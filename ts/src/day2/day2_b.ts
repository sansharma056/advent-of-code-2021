import * as fs from "fs";
import readline from "readline";

async function mulPosDepth(): Promise<number> {
	const rl = readline.createInterface(fs.createReadStream("./input"));
	let pos = 0;
	let depth = 0;
	let aim = 0;

	for await (const line of rl) {
		const instruction = line.split(" ");

		const direction = instruction[0];
		const unit = Number.parseInt(instruction[1]);

		if (direction == "forward") {
			pos += unit;
			depth += aim * unit;
		} else if (direction == "down") {
			aim += unit;
		} else if (direction == "up") {
			aim -= unit;
		}
	}

	return new Promise<number>((resolve) => resolve(pos * depth));
}

(function main() {
	mulPosDepth().then(console.log);
})();
