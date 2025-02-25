/// <reference path="../.sst/platform/config.d.ts" />


const myFunc = new sst.aws.Function( 'MyFunction', {
        handler: './lambdas',
        runtime: "go",
        url: true


        }
    );



