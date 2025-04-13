import globals from 'globals';
import pluginJs from '@eslint/js';
import tseslint from 'typescript-eslint';
import pluginVue from 'eslint-plugin-vue';
import pluginImport from 'eslint-plugin-import';
import eslintConfigPrettier from 'eslint-config-prettier';
import eslintPluginPrettierRecommended from 'eslint-plugin-prettier/recommended';
import parserVue from 'vue-eslint-parser';
import parserTs from '@typescript-eslint/parser';
import { FlatCompat } from '@eslint/eslintrc';
import path from "node:path";

const compat = new FlatCompat();

export default tseslint.config({
    files: ['**/*.{js,jsx,csj,msj}', '**/*.{ts,tsx,mts}', '**/*.vue'],
    ignores: [
        'node_modules',
        'dist',
        'public',

        'auto-imports.d.ts',
        'components.d.ts',
    ],
    extends: [
        pluginJs.configs.recommended,
        ...tseslint.configs.recommended,
        ...pluginVue.configs['flat/recommended'],
        eslintConfigPrettier,
        eslintPluginPrettierRecommended,
    ],
    plugins: {
        import: pluginImport,
    },
    languageOptions: {
        globals: globals.browser,
        parser: parserVue,
        parserOptions: {
            parser: parserTs,
            tsconfigRootDir : path.resolve('.'),
            sourceType: 'module',
        },
    },
    rules: {
        'no-shadow': 'off',
        'prefer-destructuring': 'off',

        'import/no-unresolved': 'off',
        'import/extensions': 'off',
        'import/no-absolute-path': 'off',
        'import/no-extraneous-dependencies': 'off',
        'import/order': [
            'warn',
            {
                groups: ['builtin', 'external', 'internal', 'parent', 'sibling', 'index', 'object', 'type'],
                pathGroups: [
                    {
                        pattern: '@/**',
                        group: 'internal',
                    },
                ],
                pathGroupsExcludedImportTypes: ['builtin'],
                alphabetize: {
                    order: 'asc',
                    caseInsensitive: true,
                },
                'newlines-between': 'always',
            }
        ],

        'vue/no-multiple-template-root': 'off',
        'vue/multi-word-component-names': 'off',
        'no-param-reassign': [
            'error',
            {
                props: true,
                ignorePropertyModificationsFor: ["state", "config"]
            },
        ],

        '@typescript-eslint/no-explicit-any': 'warn',
        '@typescript-eslint/no-unused-vars': 'warn',
        '@typescript-eslint/explicit-module-boundary-types': 'off',
    },
});