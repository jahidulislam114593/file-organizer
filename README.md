```markdown
# 📂 Go File Organizer

A simple and efficient CLI application written in **Golang** to automatically organize files in a directory based on their file types (extensions). Ideal for cleaning up your messy "Downloads" or "Desktop" folder with a single command!

---

## ✨ Features

- ✅ Organize files into categories like **Music**, **Videos**, **Images**, **Documents**, **Archives**, **Code**, etc.
- ✅ Automatically detects file extensions.
- ✅ Creates folders as needed.
- ✅ Cross-platform support (Linux, Windows, macOS).
- ✅ Lightweight and fast — written in Go.

---

## 📸 Example

Before:

```

Downloads/
├── song.mp3
├── movie.mkv
├── resume.pdf
├── photo.jpg
├── script.py

```

After:

```

Downloads/
├── Music/
│   └── song.mp3
├── Videos/
│   └── movie.mkv
├── Documents/
│   └── resume.pdf
├── Images/
│   └── photo.jpg
├── Code/
│   └── script.py

````

---

## 🚀 Getting Started

### 1. Clone the Repo

```bash
git clone https://github.com/yourusername/go-file-organizer.git
cd go-file-organizer
````

### 2. Build the Project

```bash
go build -o organizer
```

This will generate an executable named `organizer`.

### 3. Run the Program

```bash
./organizer
```

You'll be prompted to enter the path of the folder you want to organize.

---

## 🧠 How It Works

1. Takes a directory path as input.
2. Scans all files inside that directory (non-recursively).
3. Checks the extension of each file.
4. Moves the file into a corresponding category folder:

   * `.mp3`, `.wav` → `Music`
   * `.mp4`, `.mkv` → `Videos`
   * `.jpg`, `.png` → `Images`
   * `.pdf`, `.docx` → `Documents`
   * `.zip`, `.rar` → `Archives`
   * `.go`, `.js`, `.py` → `Code`

---

## 🛠 Technologies Used

* [Go](https://golang.org/) — Golang standard library only

---

## 📂 Supported Extensions

You can customize the file types by editing the `fileTypes` map in `main.go`.

```go
var fileTypes = map[string]string{
    ".mp3": "Music",
    ".wav": "Music",
    ".mp4": "Videos",
    ".mkv": "Videos",
    ".jpg": "Images",
    ".png": "Images",
    ".pdf": "Documents",
    ".docx": "Documents",
    ".zip": "Archives",
    ".rar": "Archives",
    ".go": "Code",
    ".js": "Code",
    ".py": "Code",
}
```

---

## ❗ Notes

* This project does **not** currently scan subdirectories (non-recursive).
* Files with unrecognized extensions are **left untouched**.
* Files are **moved**, not copied — be cautious while testing on real data.

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).

---

## 💡 Future Improvements

* Add recursive folder scanning.
* Add a dry-run or preview mode.
* Add undo support.
* Add GUI (with [Fyne](https://fyne.io/) or [Wails](https://wails.io/)).

---

## 👨‍💻 Author

**Md. Jahidul Islam**

Feel free to contribute or fork the project!

```

---

Let me know if you'd like me to add:
- Screenshots
- Badges (build, license, etc.)
- GitHub Action for automated testing
- Contribution guidelines

Or if you're publishing it to GitHub, I can generate `.gitignore` and `LICENSE` files too.
```
