import React, { useEffect, useState } from "react";
import { ListGroup } from 'react-bootstrap';
import { Objective } from "./models";
import { ObjectiveListItem } from "./ObjectiveListItem";

export const Objectives: React.FunctionComponent<{}> = () => {
    const [objectives, setObjectives] = useState<Objective[]>([]);
    const [refresh, triggerRefresh] = useState<Boolean>(true);

    useEffect(() => {
        if (!refresh) {
            return;
        }

        const apiUrl = `/api/objectives`;
        fetch(apiUrl)
            .then((res) => res.json())
            .then((objectives) => {
                setObjectives(objectives);
                triggerRefresh(false);
            });
    }, [refresh]);

    return <>
        <ListGroup>
            {objectives.map((o, id) =>
                <ObjectiveListItem key={id} objectiveId={id} objective={o} triggerRefresh={triggerRefresh}></ObjectiveListItem>
            )}
        </ListGroup>
    </>
}