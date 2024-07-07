import argparse
from Crypto.Hash import keccak

def keccak256(data: str) -> str:
    """Generate a Keccak-256 hash for the given data."""
    k = keccak.new(digest_bits=256)
    k.update(data.encode('utf-8'))
    return k.hexdigest()

def main():
    parser = argparse.ArgumentParser(description='Keccak-256 (SHA-3) Hash Generator')
    parser.add_argument('input', type=str, help='Input string to hash')
    args = parser.parse_args()

    input_data = args.input
    hash_output = keccak256(input_data)

    print(f'Input: {input_data}')
    print(f'Keccak-256 Hash: {hash_output}')

if __name__ == '__main__':
    main()