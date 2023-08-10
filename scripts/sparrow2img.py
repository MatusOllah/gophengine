#!/usr/bin/env python3

import argparse
import xml.etree.ElementTree as ET
from PIL import Image
import os
from os import path

def main():
    parser = argparse.ArgumentParser(description="unpack a sparrow v2 texture atlas")
    parser.add_argument("image", type=str, help="path to image")
    parser.add_argument("atlas", type=str, help="path to sparrow v2 texture atlas")
    parser.add_argument("-o", type=str, help="output directory", default="out/", dest="output")

    args = parser.parse_args()

    split_atlas(args.image, args.atlas, args.output)

def split_atlas(pngpath, xmlpath, output):
    if not exists(output): os.mkdir(output)

    img = Image.open(pngpath)

    print("parsing XML")
    try:
        cleaned_xml = ""
        quotepairity = 0
        with open(xmlpath, 'r', encoding='utf-8') as f:
            ch = f.read(1)
            while ch and ch != '<':
                ch = f.read(1)
            cleaned_xml += ch
            while True:
                ch = f.read(1)
                if ch == '"':
                    quotepairity = 1 - quotepairity
                elif (ch == '<' or ch == '>') and quotepairity == 1:
                    ch = '&lt;' if ch == '<' else '&gt;'
                else:
                    if not ch:
                        break
                cleaned_xml += ch

        xmltree = ET.fromstring(cleaned_xml)
    except ET.ParseError as e:
        print("error parsing XML:", str(e))
        exit(1)

    root = xmltree
    subtextures = root.findall("SubTexture")

    for subtex in subtextures:
        print(f"extracting {subtex.attrib['name']}")

        tex_x = int(subtex.attrib['x'])
        tex_y = int(subtex.attrib['y'])
        tex_width = int(subtex.attrib['width'])
        tex_height = int(subtex.attrib['height'])
        name = subtex.attrib['name']
        fx = int(subtex.attrib.get("frameX", 0))
        fy = int(subtex.attrib.get("frameY", 0))
        fw = int(subtex.attrib.get("frameWidth", tex_width))
        fh = int(subtex.attrib.get("frameHeight", tex_height))

        frame = get_true_frame(img.crop((tex_x, tex_y, tex_x+tex_width, tex_y+tex_height)).convert('RGBA'), fx, fy, fw, fh)
        frame.save(path.join(output, name+".png"))
        frame.close()

def exists(file):
    try:
        os.stat(file)
        return True
    except FileNotFoundError as e:
        return False

def get_true_frame(img, framex, framey, framew, frameh):
    final_frame = img
    if framex < 0:
        final_frame = pad_img(final_frame, False, 0, 0, 0, -framex)
    else:
        final_frame = final_frame.crop((framex, 0, final_frame.width, final_frame.height))
    
    if framey < 0:
        final_frame = pad_img(final_frame, False, -framey, 0, 0, 0)
    else:
        final_frame = final_frame.crop((0, framey, final_frame.width, final_frame.height))
    
    if framex + framew > img.width:
        final_frame = pad_img(final_frame, False, 0, framex+framew - img.width, 0, 0)
    else:
        final_frame = final_frame.crop((0, 0, framew, final_frame.height))
    
    if framey + frameh > img.height:
        final_frame = pad_img(final_frame, False, 0, 0, framey + frameh - img.height, 0)
    else:
        final_frame = final_frame.crop((0, 0, final_frame.width, frameh))

    return final_frame

def pad_img(img, clip=False, top=1, right=1, bottom=1, left=1):
    if clip:
        img = img.crop(img.getbbox())
    
    width, height = img.size
    new_width = width + right + left
    new_height = height + top + bottom
    result = Image.new('RGBA', (new_width, new_height), (0, 0, 0, 0))
    result.paste(img, (left, top))
    return result
    

if __name__ == "__main__": main()