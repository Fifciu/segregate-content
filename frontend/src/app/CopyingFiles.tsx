import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Link } from "react-router";
import { useState } from "react";
import { SelectDirectory, GetDestinationDirectory } from "../../wailsjs/go/app/DirectorySelector";

export default function CopyingFiles() {
  const [destinationDirectory, setDestinationDirectory] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  const handleSelectDestination = async () => {
    await SelectDirectory(false);
    const directory = await GetDestinationDirectory();
    setDestinationDirectory(directory);
  };

  const handleStartCopying = async () => {
    setIsLoading(true);
    try {
      // TODO: Implement copying logic here
      console.log("Starting copying to:", destinationDirectory);
      // Add your copying implementation here
    } catch (error) {
      console.error("Error during copying:", error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="container mx-auto py-8 px-4">
      <div className="mb-6">
        <h1 className="text-3xl font-bold">Kopiowanie plików</h1>
        <p className="text-muted-foreground mt-2">
          Ustaw parametry kopiowania i rozpocznij proces
        </p>
      </div>
      
      <div className="max-w-4xl">
        <Card>
          <CardHeader>
            <CardTitle>Ustawienia kopiowania</CardTitle>
            <CardDescription>
              Skonfiguruj parametry kopiowania plików
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <label className="text-sm font-medium">
                Folder docelowy
              </label>
              <Button
                type="button"
                variant="outline"
                onClick={handleSelectDestination}
                className="w-full"
                disabled={isLoading}
              >
                {destinationDirectory || 'Wybierz folder docelowy'}
              </Button>
              <p className="text-xs text-muted-foreground">
                Wybierz folder, do którego zostaną skopiowane pliki
              </p>
            </div>
            
            <div className="flex gap-4 pt-4">
              <Button asChild disabled={isLoading}>
                <Link to="/browse-project/1">Cofnij</Link>
              </Button>
              <Button 
                onClick={handleStartCopying}
                disabled={!destinationDirectory || isLoading}
              >
                {isLoading ? (
                  <div className="flex items-center gap-2">
                    <div className="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
                    Kopiowanie...
                  </div>
                ) : (
                  "Rozpocznij kopiowanie"
                )}
              </Button>
              <Button asChild disabled={isLoading}>
                <Link to="/">Home</Link>
              </Button>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
