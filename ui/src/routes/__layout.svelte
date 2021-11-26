<script lang="ts">
	import { Styles } from 'sveltestrap';
	import {
		Collapse,
		Navbar,
		NavbarToggler,
		NavbarBrand,
		Nav,
		NavItem,
		NavLink,
		Dropdown,
		DropdownToggle,
		DropdownMenu,
		DropdownItem
	} from 'sveltestrap';
	import { base } from '$app/paths';

	let isOpen = false;
	function handleUpdate(event) {
		isOpen = event.detail.isOpen;
	}
</script>

<Styles />

<Navbar color="light" light expand="md">
	<NavbarBrand href="{base}/">OKR</NavbarBrand>

	<Collapse {isOpen} navbar expand="md" on:update={handleUpdate}>
		<Nav navbar>
			<NavItem>
				<NavLink href="{base}/" />
			</NavItem>
			<NavItem>
				<NavLink href="{base}/objectives">Objectives</NavLink>
			</NavItem>
		</Nav>
	</Collapse>
</Navbar>

<main>
	<slot />
</main>
