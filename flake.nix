{
  description = "A collection of flake templates";

  outputs = { self }: {
    templates = {
      go-cli = {
        path = ./go-cli;
        description = "A simple Go cli application";
      };
    };
  };
}
