<script lang="ts" context="module">
	import { listObjectives, Objective } from '$lib/objectives';

	/**
	 * @type {import('@sveltejs/kit').Load}
	 */
	export async function load({ page, fetch, session, context }) {
		return listObjectives().then((response: Objective[]) => {
			return {
				props: {
					objectives: response
				}
			};
		});
	}
</script>

<script lang="ts">
	import ObjectiveListItem from '$lib/ObjectiveListItem.svelte';
	import { ListGroup } from 'sveltestrap';

	export let objectives: Objective[];
</script>

<ListGroup>
	{#each objectives as objective, id}
		<ObjectiveListItem objectiveId={id} {objective} />
	{/each}
</ListGroup>
