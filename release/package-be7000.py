"""Verify and copy the official static Linux ARM64 package for Xiaomi BE7000."""

import argparse
import hashlib
from pathlib import Path
import shutil
import tempfile


ARTIFACT = "sing-box-xiaomi-be7000-0.0.1.tar.gz"
UPSTREAM_ARTIFACT = "sing-box-1.14.4-linux-arm64-musl.tar.gz"
UPSTREAM_SHA256 = "b95e25f921dc5e17a4cf13327243d1177642c7e35425bf5372862c477c359f56"


def package(source, output_dir):
    output_dir.mkdir(parents=True, exist_ok=True)
    destination = output_dir / ARTIFACT
    checksum = output_dir / f"{ARTIFACT}.sha256"
    with tempfile.TemporaryDirectory(dir=output_dir) as temporary:
        staged = Path(temporary) / destination.name
        shutil.copyfile(source, staged)
        with staged.open("rb") as copied:
            digest = hashlib.file_digest(copied, "sha256").hexdigest()
        if digest != UPSTREAM_SHA256:
            raise ValueError(f"Expected the official {UPSTREAM_ARTIFACT}.")
        staged_checksum = Path(temporary) / checksum.name
        staged_checksum.write_text(f"{digest}  {destination.name}\n", encoding="ascii", newline="\n")
        staged.replace(destination)
        staged_checksum.replace(checksum)
    return destination


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("source", type=Path, help=f"Official {UPSTREAM_ARTIFACT}")
    parser.add_argument("--output-dir", type=Path, default=Path("dist/be7000"))
    args = parser.parse_args()
    try:
        print(package(args.source, args.output_dir))
    except (OSError, ValueError) as error:
        parser.exit(1, f"{parser.prog}: {error}\n")


if __name__ == "__main__":
    main()
