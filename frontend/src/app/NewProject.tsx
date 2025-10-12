import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Link } from "react-router"
import { useState } from "react"
  import { SelectFile, GetFile } from "../../wailsjs/go/app/FileSelector"
import { SelectDirectory, GetSourceDirectory } from "../../wailsjs/go/app/DirectorySelector"
import { CreateProject } from "../../wailsjs/go/app/Processor"
import { useBlocksData } from "./hooks/useBlocksData.tsx";

const HOME_COUNTRY = "Polska";

export default function NewProject() {
  const { setDays } = useBlocksData();
  const [formData, setFormData] = useState({
    name: "Islandia",
    // TODO: I could keep it in the backend to prevent overwriting
    folder: '/Volumes/Fifciuu SSD/Example',
    // folder: '/Volumes/Expansion/Zawartość/Islandia',
    planFile: null as string | null
  });
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  const handleInputChange = (field: string, value: string | null) => {
    setFormData(prev => ({
      ...prev,
      [field]: value
    }))
  }

  const handleFileSelect = async (field: string, inputType: 'file' | 'directory') => {
    if (inputType === 'file') {
      await SelectFile()
      const file = await GetFile();
      handleInputChange(field, file)
    } else {
      await SelectDirectory(true)
      const directory = await GetSourceDirectory()
      handleInputChange(field, directory)
    }
  }

  async function createProject() {
    setIsLoading(true);
    try {
      const resp = await CreateProject({
        Name: formData.name,
        Folder: formData.folder || "",
        PlanFile: formData.planFile || "",
        HomeCountry: HOME_COUNTRY
      });
      console.log(JSON.stringify(resp), 'RSP');
      setDays(resp);
    } catch (err) {
      setError(err as string)
    }
    setIsLoading(false);
  }
  return (
    <div className="container mx-auto py-8 px-4">
      <div className="mb-6">
        <h1 className="text-3xl font-bold">Nowy projekt</h1>
        <p className="text-muted-foreground mt-2">
          Utwórz nowy projekt i zacznij segregować treści
        </p>
      </div>
      
      <div className="flex gap-6 max-w-6xl">
        {/* Form Section - 60% width */}
        <div className="w-[60%]">
          <Card>
            <CardHeader>
              <CardTitle>Szczegóły projektu</CardTitle>
              <CardDescription>
                Wypełnij poniższe informacje aby utworzyć nowy projekt
              </CardDescription>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="space-y-2">
                <label htmlFor="project-name" className="text-sm font-medium">
                  Nazwa
                </label>
                <input
                  id="project-name"
                  type="text"
                  placeholder="Wprowadź nazwę projektu"
                  value={formData.name}
                  onChange={(e) => handleInputChange("name", e.target.value)}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
              </div>
              
              <div className="space-y-2">
                <label className="text-sm font-medium">
                  Folder z plikami
                </label>
                <Button
                  type="button"
                  variant="outline"
                  onClick={() => handleFileSelect("folder", "directory")}
                  className="w-full"
                >
                  {formData.folder || 'Wybierz plik'}
                </Button>
                <p className="text-xs text-muted-foreground">
                  Wybierz folder - zostanie zwrócona ścieżka bezwzględna do katalogu
                </p>
              </div>
              
              <div className="space-y-2">
                <label htmlFor="home-country" className="text-sm font-medium">
                  Kraj domowy
                </label>
                <input
                  id="home-country"
                  type="text"
                  value={HOME_COUNTRY}
                  disabled
                  className="w-full px-3 py-2 border border-gray-300 rounded-md bg-gray-100 text-gray-600 cursor-not-allowed"
                />
              </div>
              
              <div className="space-y-2">
                <label className="text-sm font-medium">
                  Plan (plik MD)
                </label>
                <Button
                  type="button"
                  variant="outline"
                  onClick={() => handleFileSelect("planFile", "file")}
                  className="w-full"
                >
                  {formData.planFile || 'Wybierz plik'}
                </Button>
                <p className="text-xs text-muted-foreground">
                  Tylko pliki z rozszerzeniem .md
                </p>
              </div>
              
              <div>
              {error && <p className="text-red-500">{error}</p>}
              <div className="flex gap-4 pt-4">
                <div className={isLoading ? "pointer-events-none opacity-50" : ""}>
                  <Button asChild>
                    <Link to="/">Anuluj</Link>
                  </Button>
                </div>
                <Button onClick={() => createProject()} disabled={isLoading}>
                  {isLoading ? (
                    <div className="flex items-center gap-2">
                      <div className="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
                      Tworzenie...
                    </div>
                  ) : (
                    "Utwórz projekt"
                  )}
                </Button>
                <div className={isLoading ? "pointer-events-none opacity-50" : ""}>
                  <Button asChild>
                    <Link to={`/browse-project/1`}>Dalej mock</Link>
                  </Button>
                </div>
              </div>
              </div>
            </CardContent>
          </Card>
        </div>

        {/* Warning Cards Section - 40% width */}
        <div className="w-[40%] space-y-4">
          <Card className="border-amber-200 bg-amber-50">
            <CardContent className="p-4">
              <p className="text-sm font-medium text-amber-800">
                Uwaga! Przygotuj klipy z różnych urządzeń w odpowiadających im katalogach.
              </p>
            </CardContent>
          </Card>

          <Card className="border-amber-200 bg-amber-50">
            <CardContent className="p-4">
              <pre className="text-sm text-amber-800 font-mono whitespace-pre">
{`Islandia/
  Lumix/
  iPhone Filip/
  iPhone Iga/
  Komarek/
  Avata/
  Insta360/`}
              </pre>
            </CardContent>
          </Card>

          <Card className="border-amber-200 bg-amber-50">
            <CardContent className="p-4">
              <p className="text-sm font-medium text-amber-800">
                Wybierz katalog "Islandia" jako "Folder z plikami"
              </p>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  )
}
