{
      description = "DevShell environment for JitRag";

      inputs = {
            nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable"; 
      };

      outputs = { self, nixpkgs }: 
            let
                  system = "x86_64-linux"; 
                  pkgs   = nixpkgs.legacyPackages.${system};
            in {
                  devShells.${system}.default = pkgs.mkShell {
                        packages = with pkgs; [
                              go
                              gopls
                              gotools
                              golangci-lint
                        ];
                        shellHook = ''
                              echo "Golang JitRag Environment"
                              go version
                              '';
                  };
            };
}
