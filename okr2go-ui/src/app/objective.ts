export class Objective {
    name: string
    keyResults: KeyResult[]
}

export class KeyResult {
    id: string
    name: string
    current: number
    target: number
    contributors: string[]
    comments: string[]
}
