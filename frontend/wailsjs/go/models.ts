export namespace main {
	
	export class Config {
	    fit_mode: string;
	    play_speed: number;
	    always_on_top: boolean;
	    window_width: number;
	    window_height: number;
	    photo_path: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fit_mode = source["fit_mode"];
	        this.play_speed = source["play_speed"];
	        this.always_on_top = source["always_on_top"];
	        this.window_width = source["window_width"];
	        this.window_height = source["window_height"];
	        this.photo_path = source["photo_path"];
	    }
	}

}

