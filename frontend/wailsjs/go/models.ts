export namespace main {
	
	export enum Weekday {
	    SUNDAY = "Sunday",
	    MONDAY = "Monday",
	    TUESDAY = "Tuesday",
	    WEDNESDAY = "Wednesday",
	    THURSDAY = "Thursday",
	    FRIDAY = "Friday",
	    SATURDAY = "Saturday",
	}
	export interface Test {
	    foo: string;
	    day: Weekday;
	}

}

export namespace test {
	
	export enum Fruit {
	    B = "Banana",
	    O = "Orange",
	}
	export enum Color {
	    G = "Green",
	    B = "Blue",
	}

}

