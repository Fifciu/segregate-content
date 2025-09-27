export namespace app {
	
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

