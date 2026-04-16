package org.example;

import java.util.Scanner;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class CodePointTest {

    @BeforeEach
    void setUp() {
        CodePoint codePoint = new CodePoint();
    }

    @Test
    void encodeUnicode() {
        String text ="he11o wor1d";

        int[] codePoints = text.codePoints().toArray();

        System.out.println(codePoints);
    }
}
