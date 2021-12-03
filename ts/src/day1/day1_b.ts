import * as fs from "fs";
import readline from "readline";

async function getDepths(): Promise<number[]> {
	const rl = readline.createInterface({
		input: fs.createReadStream("./input"),
	});

	const depths: number[] = [];

	for await (const line of rl) {
		const cur = Number.parseInt(line);
		depths.push(cur);
	}
	return new Promise<number[]>((resolve) => resolve(depths));
}

function countNumOfIncDepths(depths: number[]): number {
	let count = 0;

	let prevWindowSum = 0;
	let curWindowSum = 0;
	let windowSize = 0;

	for (let i = 0; i < depths.length; i++) {
		if (windowSize < 3) {
			windowSize++;
			curWindowSum += depths[i];
		} else {
			prevWindowSum = curWindowSum;
			curWindowSum = curWindowSum + depths[i] - depths[i - 3];

			if (curWindowSum > prevWindowSum) count++;
		}
	}

	return count;
}

(function main() {
	getDepths().then((depths) => console.log(countNumOfIncDepths(depths)));
})();
