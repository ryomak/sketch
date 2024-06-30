export interface Artwork {
    language: string;
    name: string;
}
  
export const artworks: Artwork[] = [
    { 
        language: "go", 
        name: "art_example", 
    },
    { 
        language: "go", 
        name: "ruby_image", 
    },
    { 
        language: "go", 
        name: "p5", 
    },
];

export function getArtWasm(a: Artwork) {
    return eval(getArtWasmName(a))
}

export function getArtWasmName(a: Artwork) {
    return `${a.language}_${a.name}`
}