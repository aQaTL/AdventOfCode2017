import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.stream.Collectors;
import java.util.*;

/**
 * @author aQaTL on 24.12.2017.
 */

public class Main {
	public static void main(String[] args) throws Exception {
		List<Component> components = Files.readAllLines(Paths.get("input.txt")).
				stream().
				map(s -> Arrays.stream(s.split("/")).
						mapToInt(Integer::parseInt).toArray()).
				map(ints -> new Component(ints[0], ints[1])).
				collect(Collectors.toList());

		ArrayList<ArrayList<Component>> ways = new ArrayList<>();

		components.stream().
				filter(c -> c.a == 0).
				forEach(com -> {
					matchesTo(
							new ArrayList<>(Collections.singleton(com)),
							(ArrayList<Component>) components,
							ways);
					ways.add(new ArrayList<>(Collections.singletonList(com)));
				});


		int strongest = ways.stream().
				mapToInt(way -> way.stream().
						mapToInt(com -> com.a + com.b).
						sum()).
				max().getAsInt();

		System.out.println("Part 1: " + strongest);

		int longestHighestStrength = ways.stream().
				collect(Collectors.groupingBy(ArrayList::size)).
				entrySet().
				stream().
				max(Comparator.comparing(Map.Entry::getKey)).
				get().
				getValue().
				stream().
				mapToInt(com -> com.stream().
						mapToInt(c -> c.a + c.b).
						sum()).
				max().getAsInt();

		System.out.println("Part 2: " + longestHighestStrength);
	}

	public static void matchesTo(ArrayList<Component> comps, ArrayList<Component> allComponents, ArrayList<ArrayList<Component>> allCombinations) {
		Component last = comps.get(comps.size() - 1);
		ArrayList<Component> checked = new ArrayList<>();

		allComponents.forEach(c -> {
			if (!last.equals(c) && c.a != 0 && last.matches(c) && !comps.contains(c) && !checked.contains(c)) {
				if (c.shouldSwap(last)) {
					c.swap();
				}
				checked.add(c);

				ArrayList<Component> copy = new ArrayList<>(comps);
				copy.add(c);
				allCombinations.add(copy);

				matchesTo(copy, allComponents, allCombinations);
			}
		});
	}
}

class Component {
	public int a;
	public int b;

	public Component(int a, int b) {
		this.a = a;
		this.b = b;
	}

	public void swap() {
		int tmp = a;
		a = b;
		b = tmp;
	}

	public boolean shouldSwap(Component before) {
		return before.b != this.a;
	}

	public boolean matches(Component com) {
		return this.b == com.a || this.b == com.b;
	}

	@Override
	public String toString() {
		return a + " " + b;
	}

	@Override
	public boolean equals(Object obj) {
		return obj instanceof Component &&
				((this.a == ((Component) obj).a && this.b == ((Component) obj).b) ||
						(this.a == ((Component) obj).b && this.b == ((Component) obj).a));
	}
}
