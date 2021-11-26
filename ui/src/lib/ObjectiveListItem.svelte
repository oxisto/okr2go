<script lang="ts" context="module">
</script>

<script lang="ts">
	import { Badge, Button, ListGroupItem, Progress, Table } from 'sveltestrap';
	import KeyResultRow, { KeyResultEvent } from '$lib/KeyResultRow.svelte';

	import {
		createResult,
		decreaseResultValue,
		increaseResultValue,
		KeyResult,
		Objective
	} from '$lib/objectives';

	export let objectiveId: number;
	export let objective: Objective;

	function onMinus(e: CustomEvent<KeyResultEvent>) {
		decreaseResultValue(objectiveId, e.detail.keyResult).then((result: KeyResult) => {
			objective.keyResults[e.detail.keyResultIdx] = result;
		});
	}

	function onPlus(e: CustomEvent<KeyResultEvent>) {
		increaseResultValue(objectiveId, e.detail.keyResult).then((result: KeyResult) => {
			objective.keyResults[e.detail.keyResultIdx] = result;
		});
	}

	function onNewResult(e: Event) {
		const id = prompt('Enter the id for the result', '');
		if (id == null) {
			return;
		}

		const name = prompt('Enter the name for the result', '');
		if (name == null) {
			return;
		}

		const target = prompt('Enter the target', '');
		if (target == null) {
			return;
		}

		const keyResult: KeyResult = {
			id: id,
			name: name,
			target: parseInt(target)
		};

		createResult(objectiveId, keyResult).then((keyResult: KeyResult) => {
			// update the array using an assignment
			objective.keyResults = [...objective.keyResults, keyResult];
		});
	}
</script>

<ListGroupItem>
	<div class="d-flex w-100 justify-content-between">
		<h5 class="mb-1">{objective.name}</h5>
		<div>
			<Badge color="primary" pill={true}>
				{objective.keyResults?.length} key result(s)
			</Badge>
		</div>
	</div>
	<p class="mb-1">{objective.description}</p>
	<Table class="mt-2">
		<tbody>
			{#each objective.keyResults as keyResult, keyResultIdx}
				<KeyResultRow {keyResult} {keyResultIdx} on:minus={onMinus} on:plus={onPlus} />
			{/each}
		</tbody>
	</Table>
	<Button color="primary" on:click={onNewResult}>Add Key Result</Button>
	<div>
		<small>Still working hard</small>
	</div>
</ListGroupItem>
