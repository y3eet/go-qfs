export namespace main {
	
	export class ServerInfo {
	    port: string;
	    ip: string;
	    env: string;
	    host: string;
	    os: string;
	
	    static createFrom(source: any = {}) {
	        return new ServerInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.port = source["port"];
	        this.ip = source["ip"];
	        this.env = source["env"];
	        this.host = source["host"];
	        this.os = source["os"];
	    }
	}

}

