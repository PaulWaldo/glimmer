# Project Brief: Glimmer

## Overview
Glimmer is a Flickr client application that allows users to browse their contacts' photos and photos from groups they belong to. The application provides a clean, tabbed interface for navigating between different sections and uses OAuth for authentication with Flickr.

## Core Requirements

### Authentication
- Implement OAuth-based authentication with Flickr
- Securely store user credentials in application preferences
- Provide login/logout functionality

### User Interface
- Create a main application with tabbed interface (Contacts and Groups tabs)
- Implement contact photos display with infinite scrolling
- Implement group photos display with collapsible group cards
- Create photo cards showing title, author, and image

### API Integration
- Integrate with Flickr API for authentication, contacts, and groups
- Implement methods for retrieving user's groups and photos from those groups
- Implement asynchronous image loading from Flickr URLs

### Network Visualization
- Display a visual representation of the user's Flickr network
- Show connections between people, groups, and photos
- Allow users to explore their network and discover new content

## Technical Requirements
- Follow Test-Driven Development (TDD) practices
- Ensure all operations run in the background to maintain UI responsiveness
- Implement batch loading of photos to minimize network usage
- Manage application lifecycle to ensure UI is ready before starting operations
- Make batch size and simultaneous downloads configurable

## Project Goals
- Create a responsive and user-friendly Flickr client
- Provide an intuitive way to browse and discover photos
- Reveal connections in the user's Flickr network
- Enable deeper understanding of the user's online community
