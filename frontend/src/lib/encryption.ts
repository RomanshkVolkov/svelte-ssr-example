import { Buffer } from "buffer";
import { argon2id, argon2Verify, sha256 } from "hash-wasm";

export async function hashfile(file: File) {
  if (!file) {
    throw new Error("File is required");
  }

  const arrayBuffer = await file.arrayBuffer();
  const buffer = Buffer.from(arrayBuffer);
  const hash = sha256(buffer);

  return hash;
}

export function getOTPCode() {
  // 6 digits code
  const randomBytes = new Uint8Array(3);
  crypto.getRandomValues(randomBytes);
  const code = (randomBytes[0] * 100 + randomBytes[1] * 10 + randomBytes[2]) %
    1000000;
  return code.toString().padStart(6, "0");
}

export async function hashPassword(password: string) {
  const salt = new Uint8Array(16);

  crypto.getRandomValues(salt);

  const key = await argon2id({
    password,
    salt,
    parallelism: 1,
    iterations: 256,
    memorySize: 1024,
    hashLength: 32,
    outputType: "encoded",
  });
  return key;
}

export async function verifyPassword(password: string, hash: string) {
  return await argon2Verify({
    password,
    hash,
  });
}
