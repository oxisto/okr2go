import './ObjectiveListItem.css';
import React, { SyntheticEvent } from "react"
import { Badge, Button, ListGroupItem, ProgressBar, Table } from "react-bootstrap"
import { KeyResult, Objective } from "./models"

type ObjectiveListItemProps = {
    objectiveId: number,
    objective: Objective,
    triggerRefresh: React.Dispatch<React.SetStateAction<Boolean>>
}

export const ObjectiveListItem = ({ objectiveId, objective, triggerRefresh, ...rest }: ObjectiveListItemProps) => {

    function variantForKeyResult(keyResult: KeyResult) {
        const percentage = (keyResult.current ?? 0) / keyResult.target;

        if (percentage >= 0.7) {
            return 'success';
        } else if (percentage >= 0.3) {
            return 'warning';
        } else {
            return 'danger';
        }
    }

    function onMinus(keyResult: KeyResult, e: SyntheticEvent) {
        const apiUrl = `/api/objectives/${objectiveId}/results/${keyResult.id}/minus`;
        fetch(apiUrl)
            .then((res) => res.json())
            .then((result) => {
                triggerRefresh(true);
            });
    }

    function onPlus(keyResult: KeyResult, e: SyntheticEvent) {
        const apiUrl = `/api/objectives/${objectiveId}/results/${keyResult.id}/plus`;
        fetch(apiUrl)
            .then((res) => res.json())
            .then((result) => {
                triggerRefresh(true);
            });
    }

    function onNewResult(e: SyntheticEvent) {
        const id = prompt("Enter the id for the result", "")
        if (id == null) {
            return;
        }

        const name = prompt("Enter the name for the result", "")
        if (name == null) {
            return;
        }

        const target = prompt("Enter the target", "")
        if (target == null) {
            return;
        }

        const result: KeyResult = {
            id: id,
            name: name,
            target: parseInt(target)
        };

        const apiUrl = `/api/objectives/${objectiveId}/results`;
        fetch(apiUrl, {
            method: 'POST',
            body: JSON.stringify(result),
            headers: {
                'Content-Type': 'application/json'
            }
        })
            .then((res) => res.json())
            .then((result) => {
                triggerRefresh(true);
            });
    }

    return <ListGroupItem>
        <div className="d-flex w-100 justify-content-between">
            <h5 className="mb-1">{objective.name}</h5>
            <div>
                <Badge bg="primary" pill={true}>
                    {objective.keyResults?.length} key result(s)
                </Badge>
            </div>
        </div>
        <p className="mb-1">{objective.description}</p>
        <Table className="mt-2">
            <tbody>
                {objective.keyResults?.map((keyResult, index) =>
                    <tr key={keyResult.id}>
                        <th scope="row">{keyResult.id}</th>
                        <td>{keyResult.name}</td>
                        <td>
                            <div className="d-flex">
                                <Button
                                    className="m-2"
                                    size="sm"
                                    onMouseDown={e => e.preventDefault()}
                                    onClick={e => onMinus(keyResult, e)}>-</Button>
                                <ProgressBar
                                    variant={variantForKeyResult(keyResult)}
                                    className="mt-auto mb-auto progress-bar-objective"
                                    now={keyResult.current}
                                    max={keyResult.target}
                                    label={`${keyResult.current} / ${keyResult.target}`}></ProgressBar>
                                <Button
                                    className="m-2"
                                    size="sm"
                                    onMouseDown={e => e.preventDefault()}
                                    onClick={e => onPlus(keyResult, e)}>+</Button>
                            </div>
                        </td>
                        <td>{keyResult.contributors?.join(', ')}</td>
                        <td>{keyResult.comments?.join(', ')}</td>
                    </tr >
                )}
            </tbody>
        </Table>
        <Button onClick={onNewResult}>Add Key Result</Button>
        <div>
            <small>Still working hard</small>
        </div>
    </ListGroupItem>
}

/*

/*
<tr *ngFor="let keyResult of entry.value.keyResults">
<th scope="row">{{ keyResult.id }}</th>
<td>{{ keyResult.name }}</td>
<td>
<button (click)="onMinus(entry.key, keyResult)" class="btn btn-sm btn-secondary">
<fa-icon [icon]="faMinus"></fa-icon>
</button>

<ngb-progressbar style="min-width: 300px" class="btn mt-auto mb-auto"
[type]="getTypeForKeyResult(keyResult)"
[value]="keyResult.current / keyResult.target * 100">Current
progress
is
<b>{{ keyResult.current }} /
{{ keyResult.target }}</b>...</ngb-progressbar>

<button (click)="onPlus(entry.key, keyResult)" class="btn btn-sm btn-secondary">
<fa-icon [icon]="faPlus"></fa-icon>
</button>
</td>
<td>{{ keyResult.contributors.join(', ') }}</td>
<td>{{ keyResult.comments.join(', ')}}</td>
</tr>
*/