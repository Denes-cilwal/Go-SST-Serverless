/// <reference path="./.sst/platform/config.d.ts" />

export default $config({
  app(input) {
    return {
      name: "go-sst-serverless",
      removal: input?.stage === "production" ? "retain" : "remove",
      protect: ["production"].includes(input?.stage),
      home: "aws",
      provider: {
        profile: "dinesh-serverless-test",
      }
    };
  },
  async run() {
    import("./lib/go-serverless-sst-stack");
  },
});
