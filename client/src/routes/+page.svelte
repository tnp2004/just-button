<script lang="ts">
	import { onMount } from "svelte";

	const updateEndpoint = 'http://localhost:3000/update';
	const sumEndpoint = 'http://localhost:3000/sum';
	let sum: number = 0;

	let value: number = 0;

	const updateValue = async (v: number) => {
		const res = await fetch(updateEndpoint, {
			method: 'PATCH',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
                n: v
            })
		});

        const data = await res.json()
        sum = data.n
        value = 0
	};

	const getSummation = async () => {
		const res = await fetch(sumEndpoint)
		const data = await res.json()
		sum = data.n
	}

	onMount(() => {
		// update value when render
		getSummation()
		
		const interv = setInterval(async () => await getSummation(), 5000)
		
		return () => {
			clearInterval(interv)
		}
	})


</script>

<div class="h-screen flex">
	<div class="m-auto flex flex-col gap-2 w-56">

		<div class="text-center text-emerald-500 font-bold my-10 drop-shadow-md">
			<h1 class="text-8xl">{sum}</h1>
		</div>

		<div class="mx-auto flex justify-between w-full">
			<button on:click={() => value++} class="bg-white drop-shadow-md rounded w-12 h-12 shadow-sm hover:bg-slate-100 transition duration-200 ease-linear">plus</button>
			<span class="flex flex-col justify-center text-center text-4xl font-semibold">{value}</span>
			<button on:click={() => value--} class="bg-white drop-shadow-md rounded w-12 h-12 shadow-sm hover:bg-slate-100 transition duration-200 ease-linear">minus</button>
		</div>

		<button on:click={() => updateValue(value)} class="bg-white drop-shadow-md p-2 rounded shadow-sm hover:bg-slate-100 transition duration-200 ease-linear">don't click me</button>

	</div>
</div>
