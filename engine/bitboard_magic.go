package engine

// Implement Plain Magic Bitboards
// https://www.chessprogramming.org/Magic_Bitboards#Plain

// U64 mBishopAttacks[64] [512]; // 256 K
// U64 mRookAttacks  [64][4096]; // 2048K

// struct SMagic {
//    U64 mask;  // to mask relevant squares of both lines (no outer squares)
//    U64 magic; // magic 64-bit factor
// };

// SMagic mBishopTbl[64];
// SMagic mRookTbl  [64];

// U64 bishopAttacks(U64 occ, enumSquare sq) {
//    occ &= mBishopTbl[sq].mask;
//    occ *= mBishopTbl[sq].magic;
//    occ >>= 64-9;
//    return mBishopAttacks[sq][occ];
// }
