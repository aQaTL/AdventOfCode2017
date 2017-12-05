import java.nio.file.*;
import java.util.*;

public class Main {
	public static void main(String[] args) throws Exception {
		int[] offsets = Files.readAllLines(Paths.get("input.txt")).
			stream().
			mapToInt(Integer::new).
			toArray();
		
		System.out.printf("%d\n%d", partOne(offsets.clone()), partTwo(offsets.clone()));
	}
	
	public static int partOne(int[] offsets) {
		int idx = 0;
		int steps = 0;
		
		for (;idx < offsets.length;) {
			int prev = idx;
			int off = offsets[idx];
			steps++;
			idx += off;
			offsets[prev] += 1;
		}
		return steps;
	}
	
	public static int partTwo(int[] offsets) {
		int idx = 0;
		int steps = 0;
		
		for (;idx < offsets.length;) {
			int prev = idx;
			int off = offsets[idx];
			steps++;
			idx += off;
			offsets[prev] += off >= 3 ? -1 : 1;
		}
		return steps;
	}
}
