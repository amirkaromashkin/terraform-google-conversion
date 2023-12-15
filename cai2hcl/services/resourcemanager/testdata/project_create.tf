resource "google_project" "example-project-1" {
  billing_account = "example-account-1"
  folder_id       = "456"

  labels = {
    project-label-key-a = "project-label-val-a"
  }

  name       = "My Project"
  project_id = "example-project-1"
}

resource "google_project" "example-project-2" {
  billing_account = "example-account-2"
  folder_id       = "456"

  labels = {
    project-label-key-a = "project-label-val-a"
  }

  name       = "My Project 2"
  project_id = "example-project-2"
}
