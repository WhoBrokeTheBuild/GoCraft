package main

import (
    "github.com/go-gl/mathgl/mgl32"
)

type BlockID struct {
    ID      int
    Variant int
}

type Block struct {
    Name        string
    TexCoords   []mgl32.Vec2 // Top, Bottom, Sides
}

func (b Block) GetData() ([]float32, []float32) {
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
}

func GetAllBlocks() map[BlockID]Block {
    const sq = 0.0625

    return map[BlockID]Block {
        BlockID{ 0, 0 }: Block {
            "Air", []mgl32.Vec2 {
                { -1.0, -1.0 }, { -1.0, -1.0 }, { -1.0, -1.0 }, { -1.0, -1.0 },
                { -1.0, -1.0 }, { -1.0, -1.0 }, { -1.0, -1.0 }, { -1.0, -1.0 },
                { -1.0, -1.0 }, { -1.0, -1.0 }, { -1.0, -1.0 }, { -1.0, -1.0 },
            },
        },
        BlockID{ 1, 0 }: Block {
            "Stone", []mgl32.Vec2 {
                { sq, 0.0 }, { sq * 2, 0.0 }, { sq, sq }, { sq * 2, sq },
                { sq, 0.0 }, { sq * 2, 0.0 }, { sq, sq }, { sq * 2, sq },
                { sq, 0.0 }, { sq * 2, 0.0 }, { sq, sq }, { sq * 2, sq },
            },
        },
        BlockID{ 2, 0 }: Block {
            "Grass", []mgl32.Vec2 {
                { 0.0,    0.0 }, { sq,     0.0 }, { 0.0,    sq }, { sq,     sq },
                { sq * 2, 0.0 }, { sq * 3, 0.0 }, { sq * 2, sq }, { sq * 3, sq },
                { sq * 3, 0.0 }, { sq * 4, 0.0 }, { sq * 3, sq }, { sq * 4, sq },
            },
        },
        BlockID{ 3, 0 }: Block {
            "Dirt", []mgl32.Vec2 {
                { sq * 2, 0.0 }, { sq * 3, 0.0 }, { sq * 2, sq  }, { sq * 3, sq },
                { sq * 2, 0.0 }, { sq * 3, 0.0 }, { sq * 2, sq  }, { sq * 3, sq },
                { sq * 2, 0.0 }, { sq * 3, 0.0 }, { sq * 2, sq  }, { sq * 3, sq },
            },
        },
        BlockID{ 4, 0 }: Block {
            "Cobblestone", []mgl32.Vec2 {
                { 0.0, sq }, { sq, sq }, { 0.0, sq * 2 }, { sq, sq * 2 },
                { 0.0, sq }, { sq, sq }, { 0.0, sq * 2 }, { sq, sq * 2 },
                { 0.0, sq }, { sq, sq }, { 0.0, sq * 2 }, { sq, sq * 2 },
            },
        },
    }
}
