package main

import (
    "github.com/go-gl/mathgl/mgl32"
)

type BlockShape int

const (
    ShapeCube   BlockShape = iota
    ShapeSlab   BlockShape = iota
    ShapeGrass  BlockShape = iota
)

type BlockID struct {
    ID      int
    Variant int
}

type Block struct {
    Name        string
    Shape       BlockShape
    Transparent bool
    TexCoords   []mgl32.Vec2 // Top, Bottom, Sides
}

func (b *Block) GetData() ([]float32, []float32) {
    if b.Shape == ShapeCube {
        vertices := []mgl32.Vec3 {
            // Top
            { 0.0, 1.0, 0.0 },
            { 1.0, 1.0, 0.0 },
            { 0.0, 1.0, 1.0 },
            { 1.0, 1.0, 1.0 },

            // Bottom
            { 0.0, 0.0, 0.0 },
            { 1.0, 0.0, 0.0 },
            { 0.0, 0.0, 1.0 },
            { 1.0, 0.0, 1.0 },

            // Front
            { 0.0, 1.0, 1.0 },
            { 1.0, 1.0, 1.0 },
            { 0.0, 0.0, 1.0 },
            { 1.0, 0.0, 1.0 },

            // Back
            { 0.0, 1.0, 0.0 },
        	{ 1.0, 1.0, 0.0 },
            { 0.0, 0.0, 0.0 },
        	{ 1.0, 0.0, 0.0 },

            // Left
            { 0.0, 1.0, 0.0 },
            { 0.0, 1.0, 1.0 },
            { 0.0, 0.0, 0.0 },
            { 0.0, 0.0, 1.0 },

            // Right
        	{ 1.0, 1.0, 0.0 },
            { 1.0, 1.0, 1.0 },
            { 1.0, 0.0, 0.0 },
            { 1.0, 0.0, 1.0 },
        }

        return []float32 {
            vertices[0].X(), vertices[0].Y(), vertices[0].Z(),
            vertices[1].X(), vertices[1].Y(), vertices[1].Z(),
            vertices[2].X(), vertices[2].Y(), vertices[2].Z(),
            vertices[1].X(), vertices[1].Y(), vertices[1].Z(),
            vertices[3].X(), vertices[3].Y(), vertices[3].Z(),
            vertices[2].X(), vertices[2].Y(), vertices[2].Z(),

            vertices[4].X(), vertices[4].Y(), vertices[4].Z(),
            vertices[5].X(), vertices[5].Y(), vertices[5].Z(),
            vertices[6].X(), vertices[6].Y(), vertices[6].Z(),
            vertices[6].X(), vertices[6].Y(), vertices[6].Z(),
            vertices[5].X(), vertices[5].Y(), vertices[5].Z(),
            vertices[7].X(), vertices[7].Y(), vertices[7].Z(),

            vertices[8].X(), vertices[8].Y(), vertices[8].Z(),
            vertices[9].X(), vertices[9].Y(), vertices[9].Z(),
            vertices[10].X(), vertices[10].Y(), vertices[10].Z(),
            vertices[9].X(), vertices[9].Y(), vertices[9].Z(),
            vertices[11].X(), vertices[11].Y(), vertices[11].Z(),
            vertices[10].X(), vertices[10].Y(), vertices[10].Z(),

            vertices[12].X(), vertices[12].Y(), vertices[12].Z(),
            vertices[13].X(), vertices[13].Y(), vertices[13].Z(),
            vertices[14].X(), vertices[14].Y(), vertices[14].Z(),
            vertices[14].X(), vertices[14].Y(), vertices[14].Z(),
            vertices[13].X(), vertices[13].Y(), vertices[13].Z(),
            vertices[15].X(), vertices[15].Y(), vertices[15].Z(),

            vertices[16].X(), vertices[16].Y(), vertices[16].Z(),
            vertices[17].X(), vertices[17].Y(), vertices[17].Z(),
            vertices[18].X(), vertices[18].Y(), vertices[18].Z(),
            vertices[18].X(), vertices[18].Y(), vertices[18].Z(),
            vertices[17].X(), vertices[17].Y(), vertices[17].Z(),
            vertices[19].X(), vertices[19].Y(), vertices[19].Z(),

            vertices[20].X(), vertices[20].Y(), vertices[20].Z(),
            vertices[21].X(), vertices[21].Y(), vertices[21].Z(),
            vertices[22].X(), vertices[22].Y(), vertices[22].Z(),
            vertices[22].X(), vertices[22].Y(), vertices[22].Z(),
            vertices[21].X(), vertices[21].Y(), vertices[21].Z(),
            vertices[23].X(), vertices[23].Y(), vertices[23].Z(),
        },
        []float32 {
            // Top
            b.TexCoords[0].X(), b.TexCoords[0].Y(),
            b.TexCoords[1].X(), b.TexCoords[1].Y(),
            b.TexCoords[2].X(), b.TexCoords[2].Y(),
            b.TexCoords[1].X(), b.TexCoords[1].Y(),
            b.TexCoords[3].X(), b.TexCoords[3].Y(),
            b.TexCoords[2].X(), b.TexCoords[2].Y(),

            // Bottom
            b.TexCoords[4].X(), b.TexCoords[4].Y(),
            b.TexCoords[5].X(), b.TexCoords[5].Y(),
            b.TexCoords[6].X(), b.TexCoords[6].Y(),
            b.TexCoords[6].X(), b.TexCoords[6].Y(),
            b.TexCoords[5].X(), b.TexCoords[5].Y(),
            b.TexCoords[7].X(), b.TexCoords[7].Y(),

            // Front
            b.TexCoords[8].X(), b.TexCoords[8].Y(),
            b.TexCoords[9].X(), b.TexCoords[9].Y(),
            b.TexCoords[10].X(), b.TexCoords[10].Y(),
            b.TexCoords[9].X(), b.TexCoords[9].Y(),
            b.TexCoords[11].X(), b.TexCoords[11].Y(),
            b.TexCoords[10].X(), b.TexCoords[10].Y(),

            // Back
            b.TexCoords[8].X(), b.TexCoords[8].Y(),
            b.TexCoords[9].X(), b.TexCoords[9].Y(),
            b.TexCoords[10].X(), b.TexCoords[10].Y(),
            b.TexCoords[10].X(), b.TexCoords[10].Y(),
            b.TexCoords[9].X(), b.TexCoords[9].Y(),
            b.TexCoords[11].X(), b.TexCoords[11].Y(),

            // Left
            b.TexCoords[8].X(), b.TexCoords[8].Y(),
            b.TexCoords[9].X(), b.TexCoords[9].Y(),
            b.TexCoords[10].X(), b.TexCoords[10].Y(),
            b.TexCoords[10].X(), b.TexCoords[10].Y(),
            b.TexCoords[9].X(), b.TexCoords[9].Y(),
            b.TexCoords[11].X(), b.TexCoords[11].Y(),

            // Right
            b.TexCoords[8].X(), b.TexCoords[8].Y(),
            b.TexCoords[9].X(), b.TexCoords[9].Y(),
            b.TexCoords[10].X(), b.TexCoords[10].Y(),
            b.TexCoords[10].X(), b.TexCoords[10].Y(),
            b.TexCoords[9].X(), b.TexCoords[9].Y(),
            b.TexCoords[11].X(), b.TexCoords[11].Y(),
        }
    } else if b.Shape == ShapeGrass {
        return []float32 {
            0.0, 1.0, 0.0,
            1.0, 1.0, 1.0,
            0.0, 0.0, 0.0,

            0.0, 0.0, 0.0,
            1.0, 1.0, 1.0,
            1.0, 0.0, 1.0,

            1.0, 1.0, 0.0,
            0.0, 1.0, 1.0,
            1.0, 0.0, 0.0,

            1.0, 0.0, 0.0,
            0.0, 1.0, 1.0,
            0.0, 0.0, 1.0,
        },
        []float32 {
            b.TexCoords[0].X(), b.TexCoords[0].Y(),
            b.TexCoords[1].X(), b.TexCoords[1].Y(),
            b.TexCoords[2].X(), b.TexCoords[2].Y(),

            b.TexCoords[2].X(), b.TexCoords[2].Y(),
            b.TexCoords[1].X(), b.TexCoords[1].Y(),
            b.TexCoords[3].X(), b.TexCoords[3].Y(),

            b.TexCoords[0].X(), b.TexCoords[0].Y(),
            b.TexCoords[1].X(), b.TexCoords[1].Y(),
            b.TexCoords[2].X(), b.TexCoords[2].Y(),

            b.TexCoords[2].X(), b.TexCoords[2].Y(),
            b.TexCoords[1].X(), b.TexCoords[1].Y(),
            b.TexCoords[3].X(), b.TexCoords[3].Y(),
        }
    }

    return []float32{}, []float32{}
}

func GetAllBlocks() map[BlockID]Block {
    const x = 1.0 / 32.0
    const y = 1.0 / 16.0

    return map[BlockID]Block {
        BlockID{ 0, 0 }: Block {
            "Air", ShapeCube, true, []mgl32.Vec2 {
                { -1.0, -1.0 }, { -1.0, -1.0 }, { -1.0, -1.0 }, { -1.0, -1.0 },
                { -1.0, -1.0 }, { -1.0, -1.0 }, { -1.0, -1.0 }, { -1.0, -1.0 },
                { -1.0, -1.0 }, { -1.0, -1.0 }, { -1.0, -1.0 }, { -1.0, -1.0 },
            },
        },
        BlockID{ 1, 0 }: Block {
            "Stone", ShapeCube, false, []mgl32.Vec2 {
                { x * 19, 0.0 }, { x * 20, 0.0 }, { x * 19, y }, { x * 20, y },
                { x * 19, 0.0 }, { x * 20, 0.0 }, { x * 19, y }, { x * 20, y },
                { x * 19, 0.0 }, { x * 20, 0.0 }, { x * 19, y }, { x * 20, y },
            },
        },
        BlockID{ 1, 1 }: Block {
            "Granite", ShapeCube, false, []mgl32.Vec2 {
                { x * 20, 0.0 }, { x * 21, 0.0 }, { x * 20, y }, { x * 21, y },
                { x * 20, 0.0 }, { x * 21, 0.0 }, { x * 20, y }, { x * 21, y },
                { x * 20, 0.0 }, { x * 21, 0.0 }, { x * 20, y }, { x * 21, y },
            },
        },
        BlockID{ 1, 2 }: Block {
            "Polished Granite", ShapeCube, false, []mgl32.Vec2 {
                { x * 21, 0.0 }, { x * 22, 0.0 }, { x * 21, y }, { x * 22, y },
                { x * 21, 0.0 }, { x * 22, 0.0 }, { x * 21, y }, { x * 22, y },
                { x * 21, 0.0 }, { x * 22, 0.0 }, { x * 21, y }, { x * 22, y },
            },
        },
        BlockID{ 1, 3 }: Block {
            "Diorite", ShapeCube, false, []mgl32.Vec2 {
                { x * 22, 0.0 }, { x * 23, 0.0 }, { x * 22, y }, { x * 23, y },
                { x * 22, 0.0 }, { x * 23, 0.0 }, { x * 22, y }, { x * 23, y },
                { x * 22, 0.0 }, { x * 23, 0.0 }, { x * 22, y }, { x * 23, y },
            },
        },
        BlockID{ 1, 4 }: Block {
            "Polished Diorite", ShapeCube, false, []mgl32.Vec2 {
                { x * 23, 0.0 }, { x * 24, 0.0 }, { x * 23, y }, { x * 24, y },
                { x * 23, 0.0 }, { x * 24, 0.0 }, { x * 23, y }, { x * 24, y },
                { x * 23, 0.0 }, { x * 24, 0.0 }, { x * 23, y }, { x * 24, y },
            },
        },
        BlockID{ 1, 5 }: Block {
            "Andesite", ShapeCube, false, []mgl32.Vec2 {
                { x * 24, 0.0 }, { x * 25, 0.0 }, { x * 24, y }, { x * 25, y },
                { x * 24, 0.0 }, { x * 25, 0.0 }, { x * 24, y }, { x * 25, y },
                { x * 24, 0.0 }, { x * 25, 0.0 }, { x * 24, y }, { x * 25, y },
            },
        },
        BlockID{ 1, 6 }: Block {
            "Polished Andesite", ShapeCube, false, []mgl32.Vec2 {
                { x * 25, 0.0 }, { x * 26, 0.0 }, { x * 25, y }, { x * 26, y },
                { x * 25, 0.0 }, { x * 26, 0.0 }, { x * 25, y }, { x * 26, y },
                { x * 25, 0.0 }, { x * 26, 0.0 }, { x * 25, y }, { x * 26, y },
            },
        },
        BlockID{ 2, 0 }: Block {
            "Grass", ShapeCube, false, []mgl32.Vec2 {
                { x * 2,  0.0 }, { x * 3,  0.0 }, { x * 2,  y     }, { x * 3,  y     },
                { x * 3,  0.0 }, { x * 4,  0.0 }, { x * 3,  y     }, { x * 4,  y     },
                { x * 11, y   }, { x * 12, y   }, { x * 11, y * 2 }, { x * 12, y * 2 },
            },
        },
        BlockID{ 3, 0 }: Block {
            "Dirt", ShapeCube, false, []mgl32.Vec2 {
                { x * 11, y }, { x * 12, y }, { x * 11, y * 2 }, { x * 12, y * 2 },
                { x * 11, y }, { x * 12, y }, { x * 11, y * 2 }, { x * 12, y * 2 },
                { x * 11, y }, { x * 12, y }, { x * 11, y * 2 }, { x * 12, y * 2 },
            },
        },
        BlockID{ 4, 0 }: Block {
            "Cobblestone", ShapeCube, false, []mgl32.Vec2 {
                { x * 26, 0.0 }, { x * 27, 0.0 }, { x * 26, y }, { x * 27, y },
                { x * 26, 0.0 }, { x * 27, 0.0 }, { x * 26, y }, { x * 27, y },
                { x * 26, 0.0 }, { x * 27, 0.0 }, { x * 26, y }, { x * 27, y },
            },
        },
        BlockID{ 5, 0 }: Block {
            "Oak Wood Plank", ShapeCube, false, []mgl32.Vec2 {
                { x * 14, y }, { x * 15, y }, { x * 14, y * 2 }, { x * 15, y * 2 },
                { x * 14, y }, { x * 15, y }, { x * 14, y * 2 }, { x * 15, y * 2 },
                { x * 14, y }, { x * 15, y }, { x * 14, y * 2 }, { x * 15, y * 2 },
            },
        },
        BlockID{ 5, 1 }: Block {
            "Spruce Wood Plank", ShapeCube, false, []mgl32.Vec2 {
                { x * 15, y }, { x * 16, y }, { x * 15, y * 2 }, { x * 16, y * 2 },
                { x * 15, y }, { x * 16, y }, { x * 15, y * 2 }, { x * 16, y * 2 },
                { x * 15, y }, { x * 16, y }, { x * 15, y * 2 }, { x * 16, y * 2 },
            },
        },
        BlockID{ 5, 2 }: Block {
            "Birch Wood Plank", ShapeCube, false, []mgl32.Vec2 {
                { x * 16, y }, { x * 17, y }, { x * 16, y * 2 }, { x * 17, y * 2 },
                { x * 16, y }, { x * 17, y }, { x * 16, y * 2 }, { x * 17, y * 2 },
                { x * 16, y }, { x * 17, y }, { x * 16, y * 2 }, { x * 17, y * 2 },
            },
        },
        BlockID{ 5, 3 }: Block {
            "Jungle Wood Plank", ShapeCube, false, []mgl32.Vec2 {
                { x * 17, y }, { x * 18, y }, { x * 17, y * 2 }, { x * 18, y * 2 },
                { x * 17, y }, { x * 18, y }, { x * 17, y * 2 }, { x * 18, y * 2 },
                { x * 17, y }, { x * 18, y }, { x * 17, y * 2 }, { x * 18, y * 2 },
            },
        },
        BlockID{ 5, 4 }: Block {
            "Acacia Wood Plank", ShapeCube, false, []mgl32.Vec2 {
                { x * 18, y }, { x * 19, y }, { x * 18, y * 2 }, { x * 19, y * 2 },
                { x * 18, y }, { x * 19, y }, { x * 18, y * 2 }, { x * 19, y * 2 },
                { x * 18, y }, { x * 19, y }, { x * 18, y * 2 }, { x * 19, y * 2 },
            },
        },
        BlockID{ 5, 5 }: Block {
            "Dark Oak Wood Plank", ShapeCube, false, []mgl32.Vec2 {
                { x * 19, y }, { x * 20, y }, { x * 19, y * 2 }, { x * 20, y * 2 },
                { x * 19, y }, { x * 20, y }, { x * 19, y * 2 }, { x * 20, y * 2 },
                { x * 19, y }, { x * 20, y }, { x * 19, y * 2 }, { x * 20, y * 2 },
            },
        },
        BlockID{ 6, 0 }: Block {
            "Oak Sapling", ShapeGrass, true, []mgl32.Vec2 {
                { x * 22, y * 2 }, { x * 23, y * 2 }, { x * 22, y * 3 }, { x * 23, y * 3 },
            },
        },
        BlockID{ 7, 0 }: Block {
            "Bedrock", ShapeCube, false, []mgl32.Vec2 {
                { 0.0, y }, { x, y }, { 0.0, y * 2 }, { x, y * 2 },
                { 0.0, y }, { x, y }, { 0.0, y * 2 }, { x, y * 2 },
                { 0.0, y }, { x, y }, { 0.0, y * 2 }, { x, y * 2 },
            },
        },
        BlockID{ 9, 0 }: Block {
            "Still Water", ShapeCube, false, []mgl32.Vec2 {
                { x * 26, y * 10 }, { x * 27, y * 11 }, { x * 26, y * 10 }, { x * 27, y * 11 },
                { x * 26, y * 10 }, { x * 27, y * 11 }, { x * 26, y * 10 }, { x * 27, y * 11 },
                { x * 26, y * 10 }, { x * 27, y * 11 }, { x * 26, y * 10 }, { x * 27, y * 11 },
            },
        },
        BlockID{ 17, 0 }: Block {
            "Oak Wood", ShapeCube, false, []mgl32.Vec2 {
                { x * 29, y * 2 }, { x * 30, y * 2 }, { x * 29, y * 3 }, { x * 30, y * 3 },
                { x * 29, y * 2 }, { x * 30, y * 2 }, { x * 29, y * 3 }, { x * 30, y * 3 },
                { x * 28, y * 2 }, { x * 29, y * 2 }, { x * 28, y * 3 }, { x * 29, y * 3 },
            },
        },
        BlockID{ 18, 0 }: Block {
            "Oak Leaves", ShapeCube, true, []mgl32.Vec2 {
                { x * 22, y * 4 }, { x * 23, y * 4 }, { x * 22, y * 5 }, { x * 23, y * 5 },
                { x * 22, y * 4 }, { x * 23, y * 4 }, { x * 22, y * 5 }, { x * 23, y * 5 },
                { x * 22, y * 4 }, { x * 23, y * 4 }, { x * 22, y * 5 }, { x * 23, y * 5 },
            },
        },
        BlockID{ 31, 1 }: Block {
            "Grass", ShapeGrass, true, []mgl32.Vec2 {
                { x * 18, y * 2 }, { x * 19, y * 2 }, { x * 18, y * 3 }, { x * 19, y * 3 },
            },
        },
        BlockID{ 48, 0 }: Block {
            "Moss Stone", ShapeCube, false, []mgl32.Vec2 {
                { x * 27, 0.0 }, { x * 28, 0.0 }, { x * 27, y }, { x * 28, y },
                { x * 27, 0.0 }, { x * 28, 0.0 }, { x * 27, y }, { x * 28, y },
                { x * 27, 0.0 }, { x * 28, 0.0 }, { x * 27, y }, { x * 28, y },
            },
        },
    }
}
