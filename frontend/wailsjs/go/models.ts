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

}

export namespace models {
	
	export class PreParams {
	    JavaHome: string;
	    AndroidHome: string;
	    BuildPath: string;
	    Channel: string;
	    ChannelId: string;
	    HomePath: string;
	    ExpandPath: string;
	    GamePath: string;
	    KeystoreName: string;
	    Plugins: string[];
	
	    static createFrom(source: any = {}) {
	        return new PreParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.JavaHome = source["JavaHome"];
	        this.AndroidHome = source["AndroidHome"];
	        this.BuildPath = source["BuildPath"];
	        this.Channel = source["Channel"];
	        this.ChannelId = source["ChannelId"];
	        this.HomePath = source["HomePath"];
	        this.ExpandPath = source["ExpandPath"];
	        this.GamePath = source["GamePath"];
	        this.KeystoreName = source["KeystoreName"];
	        this.Plugins = source["Plugins"];
	    }
	}

}

