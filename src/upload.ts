
document.getElementById('uploadBtn')!.addEventListener('click', async () => {
    
    const fileInput = document.getElementById('fileInput') as HTMLInputElement;
    const file: File | null = fileInput.files?.[0] ?? null;
    if(!file) {
        alert('No file selected');
        return;
    }

    const totalsize:number = file.size;
    let uploaded: number = 0;

    const stream: ReadableStream<Uint8Array> = file.stream();
    const reader: ReadableStreamDefaultReader<Uint8Array> = stream.getReader();
    
    const uploadStream = new ReadableStream<Uint8Array>({
        async pull(controller) {
            const { done, value } = await reader.read();
            if (done) {
                controller.close();
            }

            if (value) {
                uploaded += value.length;
                updateProgress(uploaded, totalsize);
            }
        }
    });
    console.log(`hi: ${file.name}`);
    await fetch(`/upload_stream?file_name=${encodeURIComponent(file.name)}`,{
        method: 'POST',
        body: uploadStream,
        duplex: 'half'
    } as RequestInit)


    alert('Upload completed')
});

function updateProgress(uploaded: number, totalsize: number) {
    const percent = (uploaded / totalsize) * 100;
    const progressBar = document.getElementById('progressBar') as HTMLElement;
    progressBar.style.width = `${percent}%`;
}
