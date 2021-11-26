export async function listObjectives(): Promise<Objective[]> {
    const apiUrl = `/api/objectives`;

    return fetch(apiUrl)
        .then((res: Response) => res.json())
        .then((response: Objective[]) => {
            return response;
        });
}

export async function decreaseResultValue(objectiveId: number, keyResult: KeyResult): Promise<KeyResult> {
    const apiUrl = `/api/objectives/${objectiveId}/results/${keyResult.id}/minus`;

    return fetch(apiUrl)
        .then((res: Response) => res.json());
}

export async function increaseResultValue(objectiveId: number, keyResult: KeyResult): Promise<KeyResult> {
    const apiUrl = `/api/objectives/${objectiveId}/results/${keyResult.id}/plus`;

    return fetch(apiUrl)
        .then((res: Response) => res.json());
}


export async function createResult(objectiveId: number, keyResult: KeyResult): Promise<KeyResult> {
    const apiUrl = `/api/objectives/${objectiveId}/results`;

    return fetch(apiUrl, {
        method: 'POST',
        body: JSON.stringify(keyResult),
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then((res) => res.json())

}

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
