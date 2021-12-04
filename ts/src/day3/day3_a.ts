import * as fs from "fs";
import readline from "readline";

async function binaryDiagnostic() {
	const rl = readline.createInterface(fs.createReadStream("./input"));

	const nums: number[][] = [];
	for await (const line of rl) {
		const row: number[] = [];
		for (const i of line) {
			row.push(Number.parseInt(i));
		}
		nums.push(row);
	}

	const counter = new Array(nums[0].length).fill(0);
	for(const row of nums) {
		for(let i=0; i<row.length; i++) {
			counter[i] += row[i];
		}
	}

	const gamma: number[] = [];
	const epsilon: number[] = [];

	for(const i of counter) {
		if(i > nums.length/2) {
			gamma.push(1);
			epsilon.push(0);
		} else {
			gamma.push(0);
			epsilon.push(1);
		}

	}

	console.log(Number.parseInt(gamma.join(""), 2) * Number.parseInt(epsilon.join(""), 2));
}

(function main() {
	binaryDiagnostic();
})();
