import { set, string} from 'fp-ts';
import * as S from 'fp-ts/string';
import { pipe } from 'fp-ts/function';

// Use bitwise combination fun to help with the translation back to represented numbers
export enum Segment {
  Top = 1,
  Middle = 2,
  Bottom = 4,
  UpperRight = 8,
  LowerRight = 16,
  UpperLeft = 32,
  LowerLeft = 64,

  One = 24,
  Two = 127 - 16 - 32,
  Three = 24 + 7,
  Four = 24 + 2 + 32,
  Five = 127 - 8 - 64,
  Six = 127 - 8,
  Seven = 24 + 1,
  Eight = 127,
  Nine = 127 - 64,
  Zero = 127 - 2,
}

export class SegmentToDecimal {
    static convertArray(segments: Segment[]): number {
        const combined = segments.reduce((acc, x) => acc | x);
        return (new SegmentToDecimal).convert(combined);
    }

    convert(segment: Segment): number {
        switch(segment) {
            case Segment.One:
                return 1;
            case Segment.Two:
                return 2;
            case Segment.Three:
                return 3;
            case Segment.Four:
                return 4;
            case Segment.Five:
                return 5;
            case Segment.Six:
                return 6;
            case Segment.Seven:
                return 7;
            case Segment.Eight:
                return 8;
            case Segment.Nine:
                return 9;
            case Segment.Zero:
                return 0;
            default:
                console.log("Unknown Segment: ", segment);
                return -1;
        }
    }
}

export class Sample {
  patterns: string[]
  outputsRaw: string[]
  lengthToCodesMap: Record<string, string[]>
  numToCodeMap: Record<string, string>
  codeToSegmentMap: Record<string, Segment>

  constructor(patterns: string[], outputs: string[]) {
    this.patterns = patterns;
    this.outputsRaw = outputs;
    this.lengthToCodesMap = {};
    this.numToCodeMap = {};
    this.codeToSegmentMap = {};
  }

  static fromLine(line: string): Sample {
    //console.log("fromLine: ", line);
    const parts = line.split("|");
    const pattern = parts[0].trim().split(" ").map(x => x.trim());
    const outputRaw = parts[1].trim().split(" ").map(x => x.trim());
    return new Sample(pattern, outputRaw);
  }

  public countEasy(): number {
    return this.outputsRaw.reduce((acc, x) => {
      return acc + ([2, 3, 4, 7].includes(x.length) ? 1 : 0);
    }, 0);
  }
}

/**
 * First find the 4 easy numbers
 * Next identify 3 and 6, which are fairly easy to find
 * Next use set math to determine what segment each code maps to
 */
export class SampleDecoder {
  public decode(sample: Sample): Sample {
    sample.lengthToCodesMap = this.buildLengthMap(sample);
    sample.numToCodeMap = this.findObviousNums(sample);
    sample.numToCodeMap = { ...sample.numToCodeMap, "3": this.findThree(sample) };
    sample.numToCodeMap = { ...sample.numToCodeMap, "6": this.findSix(sample) };
    //console.log("knowns: ", sample.numToCodeMap);
    sample.codeToSegmentMap = this.appendSegment(sample.codeToSegmentMap, this.findTop(sample), Segment.Top);
    sample.codeToSegmentMap = this.appendSegment(sample.codeToSegmentMap, this.findMiddle(sample), Segment.Middle);
    sample.codeToSegmentMap = this.appendSegment(sample.codeToSegmentMap, this.findBottom(sample), Segment.Bottom);
    sample.codeToSegmentMap = this.appendSegment(sample.codeToSegmentMap, this.findUpperRight(sample), Segment.UpperRight);
    sample.codeToSegmentMap = this.appendSegment(sample.codeToSegmentMap, this.findLowerRight(sample), Segment.LowerRight);
    sample.codeToSegmentMap = this.appendSegment(sample.codeToSegmentMap, this.findUpperLeft(sample), Segment.UpperLeft);
    sample.codeToSegmentMap = this.appendSegment(sample.codeToSegmentMap, this.findLowerLeft(sample), Segment.LowerLeft);
    //console.log("decoder: ", sample.codeToSegmentMap);
    return sample
  }

  public decodeValue(sample: Sample, value: string): number {
    const segmentCodes = value.split("");
    const segments = segmentCodes.map(x => sample.codeToSegmentMap[x]);
    return SegmentToDecimal.convertArray(segments);
  }

  public appendSegment(map: Record<string, Segment>, newCode: string, newSegment: Segment): Record<string, Segment> {
    return {
        ...map,
        [newCode]: newSegment,
    };
  }

  public findObviousNums(sample: Sample): Record<string, string> {
    return {
      "1": sample.lengthToCodesMap["2"][0],
      "7": sample.lengthToCodesMap["3"][0],
      "4": sample.lengthToCodesMap["4"][0],
      "8": sample.lengthToCodesMap["7"][0],
    }
  }

  public buildLengthMap(sample: Sample): Record<string, string[]> {
    return sample.patterns.reduce((acc, x) => {
        const len = x.length.toString();
        if (len in acc) { acc[len].push(x); } else { acc[len] = [x]; }
        return acc;
    }, {});
  }

  public findThree(sample: Sample): string {
    // find the only 5-length value that shares both of 1's codes
    const fives = sample.lengthToCodesMap[5];
    const onesCodes = sample.numToCodeMap["1"].split("")
    return fives.find(x => x.includes(onesCodes[0]) && x.includes(onesCodes[1]))
  }

  public findSix(sample: Sample): string {
    // find the only 6-length value that DOESNT share both of 1's codes
    const sixs = sample.lengthToCodesMap[6];
    const onesCodes = sample.numToCodeMap["1"].split("")
    return sixs.find(x => !x.includes(onesCodes[0]) || !x.includes(onesCodes[1]))
  }

  // top is the difference between 7 and 1
  public findTop(sample: Sample): string {
    const seven = set.fromArray(S.Eq)(sample.numToCodeMap["7"].split(""));
    const one = set.fromArray(S.Eq)(sample.numToCodeMap["1"].split(""));
    return pipe(seven, set.difference(S.Eq)(one), set.toArray(S.Ord), x => x[0]);
  }

  // lower-right is the intersection of 1 and 6
  public findLowerRight(sample: Sample): string {
    const six = set.fromArray(S.Eq)(sample.numToCodeMap["6"].split(""));
    const one = set.fromArray(S.Eq)(sample.numToCodeMap["1"].split(""));
    return pipe(six, set.intersection(S.Eq)(one), set.toArray(S.Ord), x => x[0]);
  }

  // upper-right is the component of 1 that isn't the lower-right
  public findUpperRight(sample: Sample): string {
    const lowerRight = this.findLowerRight(sample);
    return sample.numToCodeMap["1"].split("").find(x => x !== lowerRight);
  }

  // middle is the difference between (the intersection of 3 and 4) and 1
  public findMiddle(sample: Sample): string {
    const four = set.fromArray(S.Eq)(sample.numToCodeMap["4"].split(""));
    const three = set.fromArray(S.Eq)(sample.numToCodeMap["3"].split(""));
    const one = set.fromArray(S.Eq)(sample.numToCodeMap["1"].split(""));
    return pipe(four, set.intersection(S.Eq)(three), set.difference(S.Eq)(one), set.toArray(S.Ord), x => x[0]);
  }

  // upper-left is just the difference between 4 and 3
  public findUpperLeft(sample: Sample): string {
    const four = set.fromArray(S.Eq)(sample.numToCodeMap["4"].split(""));
    const three = set.fromArray(S.Eq)(sample.numToCodeMap["3"].split(""));
    return pipe(four, set.difference(S.Eq)(three), set.toArray(S.Ord), x => x[0]);
  }

  // bottom is the difference between 3 and (the union of 4 and 7)
  public findBottom(sample: Sample): string {
    const four = set.fromArray(S.Eq)(sample.numToCodeMap["4"].split(""));
    const three = set.fromArray(S.Eq)(sample.numToCodeMap["3"].split(""));
    const seven = set.fromArray(S.Eq)(sample.numToCodeMap["7"].split(""));
    const eleven = set.union(S.Eq)(seven)(four);
    return pipe(three, set.difference(S.Eq)(eleven), set.toArray(S.Ord), x => x[0]);
  }

  // lower-left can be the difference between 6 and (the union of 3 and 4)
  public findLowerLeft(sample: Sample): string {
    const four = set.fromArray(S.Eq)(sample.numToCodeMap["4"].split(""));
    const three = set.fromArray(S.Eq)(sample.numToCodeMap["3"].split(""));
    const six = set.fromArray(S.Eq)(sample.numToCodeMap["6"].split(""));
    const union = set.union(S.Eq)(three)(four);
    return pipe(six, set.difference(S.Eq)(union), set.toArray(S.Ord), x => x[0]);
  }
}
