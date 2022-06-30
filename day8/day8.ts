import * as fs from 'fs';
import {Segment, Sample, SampleDecoder, SegmentToDecimal} from './sample';

main();

function main() {
	//const inputFile = "inputSample.txt"
	const inputFile = "input.txt";
  console.log(`part1: ${part1(inputFile)}`);
  console.log(`part2: ${part2(inputFile)}`);
}

function readFileLines(filename: string): string[] {
  try {
    const fileContent = fs.readFileSync(filename, 'utf8');
    return fileContent.split("\n");
  } catch (err) {
    console.error(err);
    return [];
  }
}

function loadData(filename: string): Sample[] {
  const lines = readFileLines(filename);
  return lines.map(Sample.fromLine);
}

function part1(filename: string): number {
  const samples = loadData(filename)
  return samples.reduce((acc, x) => acc + x.countEasy(), 0);
}

function part2(filename: string): number {
  const samples = loadData(filename);
  const decoder = new SampleDecoder();
  const answers = samples.map(sample => {
    const decodedSample = decoder.decode(sample);
    const decodedOutputs = decodedSample.outputsRaw.map(x => decoder.decodeValue(decodedSample, x));
    return Number(decodedOutputs.reduce((acc, x) => acc + x.toString(), ""));
  });

  return answers.reduce((acc, x) => acc + x);
}