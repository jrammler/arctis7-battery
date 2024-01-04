{
  description = "A very simple tool to get the battery status of a Steelseries Arctis 7 headset";

  inputs.nixpkgs.url = "nixpkgs/nixos-23.11";

  outputs = { self, nixpkgs }:
  let
    nixpkgsFor = {
      x86_64-linux = import nixpkgs {
        system = "x86_64-linux";
        overlays = [ self.overlay ];
      };
    };
  in
  {
    overlay = final: prev: {

      arctis7-battery = with final; buildGoModule rec {
        pname = "arctis7-battery";
        version = "0.3.0";
        src = ./.;
        vendorHash = "sha256-0Bd8UWTohAM/RRpwuukWu9mmohG3YdLpp8bIgpEtx18=";

        buildInputs = [ hidapi udev ];
      };

    };

    packages.x86_64-linux = {
      inherit (nixpkgsFor.x86_64-linux) arctis7-battery;
    };

    defaultPackage = {
      x86_64-linux = self.packages.x86_64-linux.arctis7-battery;
    };

    nixosModule = 
    { pkgs, ... }: {
      nixpkgs.overlays = [ self.overlay ];

      environment.systemPackages = [ pkgs.arctis7-battery ];

      services.udev.extraRules = ''
            SUBSYSTEM=="hidraw", ATTRS{idVendor}=="1038", ATTRS{idProduct}=="12ad", MODE="0666"
            SUBSYSTEMS=="usb",   ATTRS{idVendor}=="1038", ATTRS{idProduct}=="12ad", MODE="0666"
      '';
    };
  };
}
