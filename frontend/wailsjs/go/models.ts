export namespace main {
	
	export class Config {
	    key: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.value = source["value"];
	    }
	}
	export class packParams {
	    JavaPath: string;
	    AndroidPath: string;
	    RootPath: string;
	    OutputPath: string;
	    ApkPath: string;
	    ProductId: string;
	    ProductName: string;
	    ApkName: string;
	    PackageName: string;
	    ChannelId: string;
	    ChannelName: string;
	
	    static createFrom(source: any = {}) {
	        return new packParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.JavaPath = source["JavaPath"];
	        this.AndroidPath = source["AndroidPath"];
	        this.RootPath = source["RootPath"];
	        this.OutputPath = source["OutputPath"];
	        this.ApkPath = source["ApkPath"];
	        this.ProductId = source["ProductId"];
	        this.ProductName = source["ProductName"];
	        this.ApkName = source["ApkName"];
	        this.PackageName = source["PackageName"];
	        this.ChannelId = source["ChannelId"];
	        this.ChannelName = source["ChannelName"];
	    }
	}

}

