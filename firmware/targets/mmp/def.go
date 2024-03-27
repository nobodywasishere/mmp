package main

import keyboard "github.com/sago35/tinygo-keyboard"

func loadKeyboardDef() {
	keyboard.KeyboardDef = []byte{
		0x5D, 0x00, 0x00, 0x80, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x3D, 0x82, 0x80, 0x17, 0x1C, 0x2E, 0x8B, 0x89, 0x9F, 0x24, 0xFF, 0xE8, 0x0F, 0x34, 0xC8, 0x2A, 0x8F, 0x45, 0xBA, 0x46, 0xF9, 0xE5, 0x77, 0x5B, 0xBF, 0x97, 0x51, 0x1F, 0xBA, 0x70, 0xAD, 0xF5, 0x40, 0xAA, 0x71, 0x76, 0x9E, 0x53, 0xDB, 0x1B, 0xD9, 0x66, 0x19, 0x94, 0xD1, 0x5A, 0xAB, 0x27, 0x7C, 0x38, 0x2E, 0xFA, 0x96, 0xE4, 0x93, 0xB2, 0x20, 0x40, 0x3D, 0x06, 0x7B, 0x31, 0x4A, 0xBD, 0x57, 0x17, 0xF3, 0xB7, 0x0A, 0x39, 0xBA, 0x7F, 0x9F, 0xCA, 0x2F, 0x9B, 0x1C, 0x05, 0x53, 0x02, 0x7F, 0x65, 0x9C, 0xF7, 0x0B, 0xDE, 0xD3, 0x51, 0xDC, 0x99, 0x3C, 0xB3, 0x64, 0xCC, 0x0C, 0x17, 0x0E, 0x58, 0x93, 0x67, 0xA1, 0xAE, 0xE7, 0xE3, 0x34, 0x08, 0x70, 0xEF, 0x31, 0x5A, 0x74, 0x70, 0x69, 0x76, 0x30, 0xC0, 0x36, 0xA4, 0x43, 0x5A, 0x2D, 0x6C, 0xAC, 0x7D, 0x0F, 0xDC, 0xC7, 0x3F, 0x8F, 0x6B, 0x78, 0xC0, 0x2F, 0xB8, 0xAB, 0xD8, 0x44, 0xBD, 0x3C, 0x96, 0x90, 0x66, 0x88, 0x4F, 0xC1, 0x3E, 0xE0, 0x2F, 0x63, 0xE3, 0x8C, 0x7E, 0xFB, 0xDD, 0x19, 0x80,
	}
}
