<script lang="ts">
	import { Button, Progress } from 'sveltestrap';
	import type { BackgroundColor } from 'sveltestrap/src/shared';
	import type { KeyResult } from './objectives';

	export let keyResult: KeyResult;
	export let keyResultIdx: number;
	export let onMinus: (keyResult: KeyResult, keyResultIdx: number, e: Event) => void;
	export let onPlus: (keyResult: KeyResult, keyResultIdx: number, e: Event) => void;

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
			<Button class="m-2" size="sm" on:click={(e) => onMinus(keyResult, keyResultIdx, e)}>-</Button>
			<div class="mb-auto mt-auto progress-bar-objective">
				<Progress
					color={colorForKeyResult(keyResult)}
					value={keyResult.current}
					max={keyResult.target}
				>
					{keyResult.current} / {keyResult.target}
				</Progress>
			</div>
			<Button class="m-2" size="sm" on:click={(e) => onPlus(keyResult, keyResultIdx, e)}>+</Button>
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
