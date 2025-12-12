export namespace tabs {
	
	export class Tab {
	    ID: string;
	    Title: string;
	    History: string[];
	    Index: number;
	    Files: string[];
	
	    static createFrom(source: any = {}) {
	        return new Tab(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Title = source["Title"];
	        this.History = source["History"];
	        this.Index = source["Index"];
	        this.Files = source["Files"];
	    }
	}

}

