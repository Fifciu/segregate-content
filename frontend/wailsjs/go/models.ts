export namespace app {
	
	export class ProcessedFile {
	    Camera?: any;
	    Filename: string;
	    CameraPath: string;
	    // Go type: time
	    NormalizedTimestamp: any;
	    // Go type: time
	    LegacyTimestamp: any;
	    Thumbnail?: string;
	
	    static createFrom(source: any = {}) {
	        return new ProcessedFile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Camera = source["Camera"];
	        this.Filename = source["Filename"];
	        this.CameraPath = source["CameraPath"];
	        this.NormalizedTimestamp = this.convertValues(source["NormalizedTimestamp"], null);
	        this.LegacyTimestamp = this.convertValues(source["LegacyTimestamp"], null);
	        this.Thumbnail = source["Thumbnail"];
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
	export class Block {
	    name: string;
	    files: ProcessedFile[];
	
	    static createFrom(source: any = {}) {
	        return new Block(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.files = this.convertValues(source["files"], ProcessedFile);
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
	export class Day {
	    blocks: Block[];
	
	    static createFrom(source: any = {}) {
	        return new Day(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.blocks = this.convertValues(source["blocks"], Block);
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
	
	export class Project {
	    Name: string;
	    Folder: string;
	    PlanFile: string;
	    HomeCountry: string;
	
	    static createFrom(source: any = {}) {
	        return new Project(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Folder = source["Folder"];
	        this.PlanFile = source["PlanFile"];
	        this.HomeCountry = source["HomeCountry"];
	    }
	}

}

export namespace maps {
	
	export class TimezoneResult {
	    dstOffset: number;
	    rawOffset: number;
	    timeZoneId: string;
	    timeZoneName: string;
	
	    static createFrom(source: any = {}) {
	        return new TimezoneResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.dstOffset = source["dstOffset"];
	        this.rawOffset = source["rawOffset"];
	        this.timeZoneId = source["timeZoneId"];
	        this.timeZoneName = source["timeZoneName"];
	    }
	}

}

