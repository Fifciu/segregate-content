import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Link } from "react-router";

export default function CopyingFiles() {
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
            <div className="text-center py-8">
              <p className="text-muted-foreground">
                Funkcjonalność kopiowania plików będzie dostępna wkrótce
              </p>
            </div>
            
            <div className="flex gap-4 pt-4">
              <Button asChild>
                <Link to="/browse-project/1">Cofnij</Link>
              </Button>
              <Button asChild>
                <Link to="/">Home</Link>
              </Button>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
