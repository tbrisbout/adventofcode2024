import { assertEquals } from "@std/assert";
import { sumDistance, similarityScore } from "./main.ts";

const testInput = `
3   4
4   3
2   5
1   3
3   9
3   3
`.trim();

Deno.test(function testSumDistance() {
  assertEquals(sumDistance(testInput), 11);
});

Deno.test(function testSimilarityScore() {
  assertEquals(similarityScore(testInput), 31);
});
