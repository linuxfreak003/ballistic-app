export namespace ports {
	
	export class Coefficient {
	    value: number;
	    name: string;
	    drag_func: number;
	
	    static createFrom(source: any = {}) {
	        return new Coefficient(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.value = source["value"];
	        this.name = source["name"];
	        this.drag_func = source["drag_func"];
	    }
	}
	export class Bullet {
	    caliber: number;
	    name: string;
	    weight: number;
	    bc?: Coefficient;
	    length: number;
	
	    static createFrom(source: any = {}) {
	        return new Bullet(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.caliber = source["caliber"];
	        this.name = source["name"];
	        this.weight = source["weight"];
	        this.bc = this.convertValues(source["bc"], Coefficient);
	        this.length = source["length"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class Environment {
	    environment_id: number;
	    temperature: number;
	    altitude: number;
	    pressure: number;
	    humidity: number;
	    wind_angle: number;
	    wind_speed: number;
	    pressure_is_absolute: boolean;
	    latitude: number;
	    azimuth: number;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Environment(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.environment_id = source["environment_id"];
	        this.temperature = source["temperature"];
	        this.altitude = source["altitude"];
	        this.pressure = source["pressure"];
	        this.humidity = source["humidity"];
	        this.wind_angle = source["wind_angle"];
	        this.wind_speed = source["wind_speed"];
	        this.pressure_is_absolute = source["pressure_is_absolute"];
	        this.latitude = source["latitude"];
	        this.azimuth = source["azimuth"];
	        this.name = source["name"];
	    }
	}
	export class Load {
	    load_id: number;
	    bullet?: Bullet;
	    muzzle_velocity: number;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Load(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.load_id = source["load_id"];
	        this.bullet = this.convertValues(source["bullet"], Bullet);
	        this.muzzle_velocity = source["muzzle_velocity"];
	        this.name = source["name"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Rifle {
	    RifleId: number;
	    name: string;
	    sight_height: number;
	    barrel_twist: number;
	    twist_direction_left: boolean;
	    zero_range: number;
	
	    static createFrom(source: any = {}) {
	        return new Rifle(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.RifleId = source["RifleId"];
	        this.name = source["name"];
	        this.sight_height = source["sight_height"];
	        this.barrel_twist = source["barrel_twist"];
	        this.twist_direction_left = source["twist_direction_left"];
	        this.zero_range = source["zero_range"];
	    }
	}
	export class Scenario {
	    scenario_id: number;
	    name: string;
	    environment_id: number;
	    rifle_id: number;
	    load_id: number;
	
	    static createFrom(source: any = {}) {
	        return new Scenario(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.scenario_id = source["scenario_id"];
	        this.name = source["name"];
	        this.environment_id = source["environment_id"];
	        this.rifle_id = source["rifle_id"];
	        this.load_id = source["load_id"];
	    }
	}

}

