export namespace main {
	
	export class Config {
	    fit_mode: string;
	    play_speed: number;
	    always_on_top: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fit_mode = source["fit_mode"];
	        this.play_speed = source["play_speed"];
	        this.always_on_top = source["always_on_top"];
	    }
	}

}

