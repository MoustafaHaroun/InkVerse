export async function load({params, fetch}) {

    async function fetchNovel(id: string) {
        const novelRes = await fetch('http://localhost:8000/novels/' + id)
        const novelData = await novelRes.json();

        return novelData;
    }

    async function fetchChapters(id: string) {
        const url : string = 'http://localhost:8000/novels/' + id + '/chapters'

        const chapterRes = await fetch(url)
        const chapterData : any[] = await chapterRes.json()

        return chapterData; 
    }

    return{
        novel: await fetchNovel(params.novelId),
        chapters: await fetchChapters(params.novelId)
    }
}
