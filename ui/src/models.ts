export interface Objective {
    name: string
    description?: string
    keyResults?: KeyResult[]
}

export interface KeyResult {
    id: string,
    name: string
    current?: number,
    target: number,
    contributors?: string[],
    comments?: string[]
}
