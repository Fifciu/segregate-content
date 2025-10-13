import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Link } from "react-router";
import { CheckCircle2 } from "lucide-react";

export default function SucceedPage() {
  return (
    <div className="container mx-auto py-8 px-4">
      <div className="mb-6">
        <h1 className="text-3xl font-bold">Sukces!</h1>
        <p className="text-muted-foreground mt-2">
          Pliki zostały pomyślnie skopiowane
        </p>
      </div>
      
      <div className="max-w-4xl">
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <CheckCircle2 className="w-8 h-8 text-green-500" />
              Operacja zakończona sukcesem
            </CardTitle>
            <CardDescription>
              Wszystkie pliki zostały prawidłowo skopiowane do wybranego folderu
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="bg-green-50 border border-green-200 rounded-lg p-4">
              <p className="text-sm text-green-800">
                Pliki zostały zorganizowane według struktury dni i pomyślnie przeniesione do wskazanej lokalizacji.
              </p>
            </div>
            
            <div className="flex gap-4 pt-4">
              <Button asChild>
                <Link to="/">Powrót do strony głównej</Link>
              </Button>
              <Button asChild variant="outline">
                <Link to="/new-project">Nowy projekt</Link>
              </Button>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}

