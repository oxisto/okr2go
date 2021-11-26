<script lang="ts" context="module">
	export interface KeyResultEvent {
		keyResult: KeyResult;
		keyResultIdx: number;
	}

	export interface KeyResultEventMap {
		minus: KeyResultEvent;
		plus: KeyResultEvent;
	}
</script>

<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	import { Button, Progress } from 'sveltestrap';
	import type { BackgroundColor } from 'sveltestrap/src/shared';
	import type { KeyResult } from './objectives';

	export let keyResult: KeyResult;
	export let keyResultIdx: number;

	const dispatch = createEventDispatcher<KeyResultEventMap>();

	function colorForKeyResult(keyResult: KeyResult): BackgroundColor {
		const percentage = (keyResult.current ?? 0) / keyResult.target;

		if (percentage >= 0.7) {
			return 'success';
		} else if (percentage >= 0.3) {
			return 'warning';
		} else {
			return 'danger';
		}
	}
</script>

<tr>
	<th scope="row">{keyResult.id}</th>
	<td>{keyResult.name}</td>
	<td>
		<div class="d-flex">
			<Button
				class="m-2"
				size="sm"
				on:click={(e) => dispatch('minus', { keyResult: keyResult, keyResultIdx: keyResultIdx })}
			>
				-
			</Button>
			<div class="mb-auto mt-auto progress-bar-objective">
				<Progress
					color={colorForKeyResult(keyResult)}
					value={keyResult.current}
					max={keyResult.target}
				>
					{keyResult.current} / {keyResult.target}
				</Progress>
			</div>
			<Button
				class="m-2"
				size="sm"
				on:click={(e) => dispatch('plus', { keyResult: keyResult, keyResultIdx: keyResultIdx })}
			>
				+
			</Button>
		</div>
	</td>
	<td>{keyResult.contributors?.join(', ')}</td>
	<td>{keyResult.comments?.join(', ')}</td>
</tr>

<style>
	.progress-bar-objective {
		min-width: 10rem;
	}
</style>
