import { seedUsers } from "./users";

async function mainSeed() {
  await Promise.all(
    [
      seedUsers(),
    ],
  );
}

(async () => mainSeed())();
