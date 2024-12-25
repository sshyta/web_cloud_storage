async function fetchFiles() {
    const response = await fetch('/files');
    const files = await response.json();
    const fileList = document.getElementById('fileList');
    fileList.innerHTML = '';

    files.forEach(file => {
        const fileItem = document.createElement('div');
        fileItem.classList.add('file-item');
        fileItem.innerHTML = `
                <span class="file-name">${file.name}</span>
                <div class="file-actions">
                    <button class="view-btn" onclick="viewFile('${file.name}')">View</button>
                    <button class="edit-btn" onclick="editFile('${file.name}')">Edit</button>
                    <button class="delete-btn" onclick="deleteFile('${file.name}')">Delete</button>
                </div>
            `;
        fileList.appendChild(fileItem);
    });
}

function searchFiles() {
    const query = document.getElementById('searchBox').value.toLowerCase();
    const fileItems = document.querySelectorAll('.file-item');
    fileItems.forEach(item => {
        const fileName = item.querySelector('.file-name').textContent.toLowerCase();
        item.style.display = fileName.includes(query) ? '' : 'none';
    });
}

async function viewFile(fileName) {
    window.location.href = `/view?file=${fileName}`;
}

async function editFile(fileName) {
    window.location.href = `/edit?file=${fileName}`;
}

async function deleteFile(fileName) {
    if (confirm(`Are you sure you want to delete ${fileName}?`)) {
        const response = await fetch(`/delete?file=${fileName}`, { method: 'DELETE' });
        if (response.ok) {
            alert('File deleted successfully');
            fetchFiles();
        } else {
            alert('Failed to delete file');
        }
    }
}