import { input } from "./input.ts";

const parseInput = (input: string): [number[], number[]] => 
  input.split("\n")
    .reduce<[number[], number[]]>((acc, line) => {
      const [l, r] = line.split(/\s+/)
      acc[0].push(parseInt(l, 10));
      acc[1].push(parseInt(r, 10));
      return acc;
    } , [[], []]);

export const sumDistance = (input: string): number => {
  const [left, right] = parseInput(input);

  const l = left.sort()
  const r = right.sort()

  return l.reduce((sum, a, i) => sum + Math.abs(a - r[i]), 0)
}

export const similarityScore = (input: string): number => {
  const [left, right] = parseInput(input);

  return left.reduce((score, a) => score + (a * right.filter(b => a === b).length), 0)
}

// Learn more at https://docs.deno.com/runtime/manual/examples/module_metadata#concepts
if (import.meta.main) {
  console.log("Part 1:", sumDistance(input))
  console.log("Part 2:", similarityScore(input))
}
