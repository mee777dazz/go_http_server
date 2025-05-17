import os
import pytest
import requests

@pytest.fixture
def base_url():
    return os.getenv("BASE_URL", "http://localhost:8080")

def test_post_object(base_url):
    data = {"key": "testKey", "value": "testValue"}
    response = requests.post(f"{base_url}/object", json=data)
    assert response.status_code == 200

def test_get_object(base_url):
    response = requests.get(f"{base_url}/object?key=testKey")
    assert response.status_code == 200
    assert response.json()["value"] == "testValue"

def test_put_object(base_url):
    data = {"key": "testKey", "value": "updatedValue"}
    response = requests.put(f"{base_url}/object", json=data)
    assert response.status_code == 200

def test_get_updated_object(base_url):
    response = requests.get(f"{base_url}/object?key=testKey")
    assert response.status_code == 200
    assert response.json()["value"] == "updatedValue"

def test_delete_object(base_url):
    response = requests.delete(f"{base_url}/object?key=testKey")
    assert response.status_code == 200

def test_get_deleted_object(base_url):
    response = requests.get(f"{base_url}/object?key=testKey")
    assert response.status_code == 404