<script lang="ts">
    import { onMount } from "svelte";

    type Novel = {
        id: string;
        author_id: string;
        title: string;
        synopsis: string;
        rating: number;
        created_at: string;
    }

    let novels: Novel[] = [];

    async function fetchNovels() {
        await fetch('http://localhost:8000/novels/')
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json(); 
        })
        .then(data => novels = data )
        .then(data => console.log(data))
        .catch(error => console.log(error));
    }

    onMount(fetchNovels)
</script>

<main class="container mx-auto px-4 py-8 mt-20">

    <table>
        <thead class="table-auto border bg-gray-800 text-white">
            <th class="rounded-tl-md">Series</th>
            <th>Synopsis</th>
            <th class="rounded-br-md">Rating</th>
        </thead>
        <tbody>
            {#each novels as novel }
            <tr>
                <a href="/novel/{novel.id}"><td>{novel.title}</td></a>
                <td>{novel.synopsis}</td>
                <td>{novel.rating}</td>
            </tr>
            {/each}
        </tbody>
    </table>
</main>

