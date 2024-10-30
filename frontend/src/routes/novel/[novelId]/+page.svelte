<script lang="ts">
    export let data;

</script>

<main class="container mx-auto px-4 py-8 mt-20">
    <h1>Novel Page</h1>
    <img src="https://cdn.scribblehub.com/images/36/hell-difficulty-tutorial_727602_1703253381.jpg" alt="novel cover">
    <h1>{data.novel.title}</h1>
    <h1>{data.novel.synopsis}</h1>
    <h1>{data.novel.rating}</h1>

    <table>
        <thead class="table-auto border bg-gray-800 text-white">
            <th class="rounded-tl-md">Title</th>
        </thead>
        <tbody>
            {#each data.chapters as chapter }
            <tr>
                <a href="/chapter/{chapter.id}"><td>{chapter.title}</td></a>
            </tr>
            {/each}
        </tbody>
    </table>
</main>