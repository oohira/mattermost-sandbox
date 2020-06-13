serve:
	docker run --rm --name mattermost-preview -d -p 8065:8065 --add-host dockerhost:127.0.0.1 mattermost/mattermost-preview
